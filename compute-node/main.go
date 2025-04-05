package main

import (
	"encoding/hex"
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"gocv.io/x/gocv"

	// This would be your abigen-generated package
	"github.com/Ethglobal-taipei/Admojo-module/compute-node/oracle"
)

// ---------------------------------------------------------------------
// A. Data Structures
// ---------------------------------------------------------------------
type DeviceConfig struct {
	StreamURL  string
	DeviceID   uint64
	TSWriteKey string
}

// ---------------------------------------------------------------------
// B. Main Program
// ---------------------------------------------------------------------
func main() {
	// Load environment
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found or could not load.")
	}

	// Setup intervals
	aggSeconds, _ := strconv.Atoi(getEnv("AGGREGATION_INTERVAL", "300"))
	aggregationInterval := time.Duration(aggSeconds) * time.Second

	frameSeconds, _ := strconv.Atoi(getEnv("FRAME_INTERVAL", "5"))
	frameInterval := time.Duration(frameSeconds) * time.Second

	// Setup DNN thresholds
	dnnConfidenceStr := getEnv("DNN_CONFIDENCE_THRESHOLD", "0.5")
	dnnConfidence, _ := strconv.ParseFloat(dnnConfidenceStr, 64)

	dupThresholdStr := getEnv("DUPLICATE_IOU_THRESHOLD", "0.5")
	dupThreshold, _ := strconv.ParseFloat(dupThresholdStr, 64)

	// Setup Ethereum (PerformanceOracle)
	transactor, perfOracle := setupEthereum()

	// Load DNN model
	protoPath := getEnv("DNN_PROTO_PATH", "./deploy.prototxt")
	modelPath := getEnv("DNN_MODEL_PATH", "./res10_300x300_ssd_iter_140000.caffemodel")
	net := gocv.ReadNetFromCaffe(protoPath, modelPath)
	if net.Empty() {
		log.Fatalf("Error loading DNN model from %s / %s\n", protoPath, modelPath)
	}
	defer net.Close()

	// Gather device configs from .env
	devices := readDeviceConfigs()

	// Start a goroutine for each device
	var wg sync.WaitGroup
	wg.Add(len(devices))
	for _, cfg := range devices {
		go func(dc DeviceConfig) {
			defer wg.Done()
			runDevice(dc, &net, transactor, perfOracle,
				aggregationInterval,
				frameInterval,
				float32(dnnConfidence),
				float32(dupThreshold),
			)
		}(cfg)
	}

	log.Println("All device loops started. Press Ctrl+C to stop.")
	wg.Wait()
	log.Println("All device loops ended. Exiting.")
}

// ---------------------------------------------------------------------
// C. runDevice
// ---------------------------------------------------------------------
func runDevice(
	cfg DeviceConfig,
	net *gocv.Net,
	transactor *bind.TransactOpts,
	perfOracle *oracle.PerformanceOracle,
	aggregationInterval time.Duration,
	frameInterval time.Duration,
	dnnConfidenceThreshold float32,
	duplicateThreshold float32,
) {
	// 1) Do an HTTP GET to read firmware headers (X-Firmware-Hash, X-Signature)
	firmwareHash, firmwareSig, err := fetchFirmwareHeaders(cfg.StreamURL)
	if err != nil {
		log.Printf("[Device %d] WARNING: Could not fetch firmware headers: %v", cfg.DeviceID, err)
		// You might decide to return or keep going. We'll proceed, but the
		// on-chain call won't work if these are blank.
	} else {
		log.Printf("[Device %d] Fetched FirmwareHash=%s, Signature=%s\n", cfg.DeviceID, firmwareHash, firmwareSig)
	}

	// 2) Open the MJPEG stream
	cap, err := gocv.OpenVideoCapture(cfg.StreamURL)
	if err != nil {
		log.Printf("[Device %d] ERROR: Failed to open stream %s: %v\n", cfg.DeviceID, cfg.StreamURL, err)
		return
	}
	defer cap.Close()

	window := gocv.NewWindow(fmt.Sprintf("Device-%d", cfg.DeviceID))
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	aggregatedCount := 0
	startTime := time.Now()

	var previousDetections []image.Rectangle
	log.Printf("[Device %d] Starting detection loop on %s\n", cfg.DeviceID, cfg.StreamURL)

	for {
		if ok := cap.Read(&img); !ok {
			log.Printf("[Device %d] Cannot read frame, stopping.\n", cfg.DeviceID)
			break
		}
		if img.Empty() {
			time.Sleep(frameInterval)
			continue
		}

		// Convert frame to a 300x300 blob
		blob := gocv.BlobFromImage(img, 1.0, image.Pt(300, 300),
			gocv.NewScalar(104.0, 177.0, 123.0, 0), false, false)
		net.SetInput(blob, "")
		detections := net.Forward("")
		blob.Close()

		totalDetections := int(detections.Total() / 7)
		detectionMat := detections.Reshape(1, totalDetections)

		currentDetections := make([]image.Rectangle, 0, totalDetections)
		newFaces := 0

		for i := 0; i < totalDetections; i++ {
			confidence := detectionMat.GetFloatAt(i, 2)
			if confidence < dnnConfidenceThreshold {
				continue
			}
			// bounding box
			left := int(detectionMat.GetFloatAt(i, 3) * float32(img.Cols()))
			top := int(detectionMat.GetFloatAt(i, 4) * float32(img.Rows()))
			right := int(detectionMat.GetFloatAt(i, 5) * float32(img.Cols()))
			bottom := int(detectionMat.GetFloatAt(i, 6) * float32(img.Rows()))
			rect := image.Rect(left, top, right, bottom)
			currentDetections = append(currentDetections, rect)

			// check for duplication
			duplicate := false
			for _, prevRect := range previousDetections {
				if iou(rect, prevRect) >= float64(duplicateThreshold) {
					duplicate = true
					break
				}
			}
			if !duplicate {
				newFaces++
			}

			// draw bounding box
			gocv.Rectangle(&img, rect, color.RGBA{0, 255, 0, 0}, 2)
			label := fmt.Sprintf("%.2f", confidence)
			size := gocv.GetTextSize(label, gocv.FontHersheySimplex, 0.5, 1)
			pt := image.Pt(left, top-2)
			gocv.Rectangle(&img, image.Rect(pt.X, pt.Y-size.Y, pt.X+size.X, pt.Y),
				color.RGBA{0, 255, 0, 0}, -1)
			gocv.PutText(&img, label, pt, gocv.FontHersheySimplex, 0.5,
				color.RGBA{0, 0, 0, 0}, 1)
		}

		aggregatedCount += newFaces
		previousDetections = currentDetections
		log.Printf("[Device %d] Detected %d new face(s) this frame, total so far: %d\n", cfg.DeviceID, newFaces, aggregatedCount)

		// Show frame
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}

		// Time to push data?
		if time.Since(startTime) >= aggregationInterval {
			log.Printf("[Device %d] Aggregation interval reached. Count=%d\n", cfg.DeviceID, aggregatedCount)

			// 1) Send to ThingSpeak
			errTS := sendDataToThingSpeak(
				"https://api.thingspeak.com/update",
				cfg.TSWriteKey,
				aggregatedCount,
			)
			if errTS != nil {
				log.Printf("[Device %d] Error pushing to ThingSpeak: %v\n", cfg.DeviceID, errTS)
			} else {
				log.Printf("[Device %d] Successfully sent to ThingSpeak.\n", cfg.DeviceID)
			}

			// 2) Send on-chain with signature
			timestamp := uint64(time.Now().Unix())

			// convert the device’s firmwareHash and signature from hex to raw bytes
			fwBytes, errFh := hex.DecodeString(strings.TrimPrefix(firmwareHash, "0x"))
			sigBytes, errSig := hex.DecodeString(strings.TrimPrefix(firmwareSig, "0x"))
			if errFh != nil || errSig != nil {
				log.Printf("[Device %d] Invalid hex firmwareHash or signature: %v, %v\n", cfg.DeviceID, errFh, errSig)
			} else {
				// Ensure 32 bytes for fwHash
				if len(fwBytes) == 32 && len(sigBytes) == 65 {
					var fwArr [32]byte
					copy(fwArr[:], fwBytes)

					tx, errTx := perfOracle.UpdateMetricsWithSig(
						transactor,
						big.NewInt(int64(cfg.DeviceID)),
						big.NewInt(int64(timestamp)),
						big.NewInt(int64(aggregatedCount)),
						big.NewInt(0), // taps=0 for example
						fwArr,
						sigBytes,
					)
					if errTx != nil {
						log.Printf("[Device %d] updateMetricsWithSig error: %v\n", cfg.DeviceID, errTx)
					} else {
						log.Printf("[Device %d] updateMetricsWithSig TX: %s\n", cfg.DeviceID, tx.Hash().Hex())
					}
				} else {
					log.Printf("[Device %d] FirmwareHash or signature has incorrect length!\n", cfg.DeviceID)
				}
			}

			// Reset
			aggregatedCount = 0
			startTime = time.Now()
			previousDetections = nil
		}

		time.Sleep(frameInterval)
	}
}

// ---------------------------------------------------------------------
// D. Support Functions
// ---------------------------------------------------------------------

// readDeviceConfigs reads environment variables for device1, device2, etc.
func readDeviceConfigs() []DeviceConfig {
	dev1URL := getEnv("ESP32_STREAM_URL1", "")
	dev1IDStr := getEnv("ESP32_STREAM_ID1", "0")
	dev1TSKey := getEnv("THINGSPEAK_WRITE_KEY1", "")

	dev2URL := getEnv("ESP32_STREAM_URL2", "")
	dev2IDStr := getEnv("ESP32_STREAM_ID2", "0")
	dev2TSKey := getEnv("THINGSPEAK_WRITE_KEY2", "")

	d1ID, _ := strconv.ParseUint(dev1IDStr, 10, 64)
	d2ID, _ := strconv.ParseUint(dev2IDStr, 10, 64)

	devices := []DeviceConfig{
		{
			StreamURL:  dev1URL,
			DeviceID:   d1ID,
			TSWriteKey: dev1TSKey,
		},
		{
			StreamURL:  dev2URL,
			DeviceID:   d2ID,
			TSWriteKey: dev2TSKey,
		},
	}
	return devices
}

// setupEthereum sets up an ethclient, parses the private key, creates a transactor,
// and binds to the PerformanceOracle contract.
func setupEthereum() (*bind.TransactOpts, *oracle.PerformanceOracle) {
	rpcURL := getEnv("RPC_URL", "")
	contractAddrStr := getEnv("CONTRACT_ADDRESS", "")
	privKeyHex := getEnv("ADMIN_PRIVATE_KEY", "")
	chainIDStr := getEnv("CHAIN_ID", "5")

	chainID, err := strconv.ParseInt(chainIDStr, 10, 64)
	if err != nil {
		log.Fatalf("Invalid CHAIN_ID: %v", err)
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Error creating ETH client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privKeyHex, "0x"))
	if err != nil {
		log.Fatalf("Invalid ADMIN_PRIVATE_KEY: %v", err)
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		log.Fatalf("Failed to create keyed transactor: %v", err)
	}

	// bind the contract
	contractAddress := common.HexToAddress(strings.TrimPrefix(contractAddrStr, "0x"))
	perfOracle, err := oracle.NewPerformanceOracle(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to bind PerformanceOracle at %s: %v", contractAddrStr, err)
	}

	return transactor, perfOracle
}

// fetchFirmwareHeaders does a quick GET (or HEAD) to the device URL
// to retrieve the X-Firmware-Hash and X-Signature headers.
func fetchFirmwareHeaders(streamURL string) (string, string, error) {
	req, err := http.NewRequest("GET", streamURL, nil)
	if err != nil {
		return "", "", err
	}
	// We only need headers, but let's do a GET and discard the body
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	// read up to 1KB just so we fully consume the response
	io.CopyN(io.Discard, resp.Body, 1024)

	fwHash := resp.Header.Get("X-Firmware-Hash")
	fwSig := resp.Header.Get("X-Signature")
	if fwHash == "" || fwSig == "" {
		return fwHash, fwSig, fmt.Errorf("missing X-Firmware-Hash or X-Signature")
	}

	return fwHash, fwSig, nil
}

// sendDataToThingSpeak posts the aggregated face count to ThingSpeak.
func sendDataToThingSpeak(tsURL, tsKey string, faceCount int) error {
	endpoint := fmt.Sprintf("%s?api_key=%s&field1=%d", tsURL, tsKey, faceCount)
	resp, err := http.Post(endpoint, "application/x-www-form-urlencoded", nil)
	if err != nil {
		return fmt.Errorf("HTTP POST request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("ThingSpeak responded with status: %s", resp.Status)
	}
	return nil
}

// iou calculates Intersection over Union of two rectangles
func iou(a, b image.Rectangle) float64 {
	intersect := a.Intersect(b)
	if intersect.Empty() {
		return 0.0
	}
	interArea := float64(intersect.Dx() * intersect.Dy())
	areaA := float64(a.Dx() * a.Dy())
	areaB := float64(b.Dx() * b.Dy())
	unionArea := areaA + areaB - interArea
	return interArea / unionArea
}

// getEnv returns the env variable key if present, else fallback
func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

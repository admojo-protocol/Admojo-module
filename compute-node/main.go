package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gocv.io/x/gocv"
)

// IoU computes the Intersection over Union of two rectangles.
func IoU(a, b image.Rectangle) float64 {
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

func main() {
	// ------------------------------------------------
	// 1. Load configuration from .env
	// ------------------------------------------------
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: Could not load .env file.")
	}

	streamURL := getEnv("ESP32_STREAM_URL", "http://192.168.1.43/")
	dnnProtoPath := getEnv("DNN_PROTO_PATH", "./deploy.prototxt")
	dnnModelPath := getEnv("DNN_MODEL_PATH", "./res10_300x300_ssd_iter_140000.caffemodel")
	dnnConfidenceThresholdStr := getEnv("DNN_CONFIDENCE_THRESHOLD", "0.5")
	dupThresholdStr := getEnv("DUPLICATE_IOU_THRESHOLD", "0.5")
	thingspeakWriteKey := getEnv("THINGSPEAK_WRITE_KEY", "")
	thingspeakURL := getEnv("THINGSPEAK_URL", "https://api.thingspeak.com/update")

	aggIntervalSeconds := getEnv("AGGREGATION_INTERVAL", "300")
	aggInterval, err := strconv.Atoi(aggIntervalSeconds)
	if err != nil {
		aggInterval = 300
	}
	aggregationInterval := time.Duration(aggInterval) * time.Second

	frameIntervalSeconds := getEnv("FRAME_INTERVAL", "5")
	frameInt, err := strconv.Atoi(frameIntervalSeconds)
	if err != nil {
		frameInt = 5
	}
	frameInterval := time.Duration(frameInt) * time.Second

	if thingspeakWriteKey == "" {
		log.Fatal("THINGSPEAK_WRITE_KEY not set. Please configure .env properly.")
	}

	dnnConfidenceThreshold, err := strconv.ParseFloat(dnnConfidenceThresholdStr, 64)
	if err != nil {
		dnnConfidenceThreshold = 0.5
	}

	duplicateThreshold, err := strconv.ParseFloat(dupThresholdStr, 64)
	if err != nil {
		duplicateThreshold = 0.5
	}

	// ------------------------------------------------
	// 2. Open the ESP32-CAM MJPEG stream
	// ------------------------------------------------
	cap, err := gocv.OpenVideoCapture(streamURL)
	if err != nil {
		log.Fatalf("Error opening video capture from %s: %v", streamURL, err)
	}
	defer cap.Close()

	// ------------------------------------------------
	// 3. Create a window for visual debugging
	// ------------------------------------------------
	window := gocv.NewWindow("DNN Face Detection")
	defer window.Close()

	// ------------------------------------------------
	// 4. Load the deep learning face detection model
	// ------------------------------------------------
	net := gocv.ReadNetFromCaffe(dnnProtoPath, dnnModelPath)
	if net.Empty() {
		log.Fatalf("Error reading network model from: %s and %s", dnnProtoPath, dnnModelPath)
	}
	defer net.Close()

	// ------------------------------------------------
	// 5. Prepare a Mat to store frames
	// ------------------------------------------------
	img := gocv.NewMat()
	defer img.Close()

	// previousDetections will hold the bounding boxes from the last frame.
	var previousDetections []image.Rectangle

	aggregatedCount := 0
	startTime := time.Now()

	log.Printf("Starting DNN-based face detection. Capturing one frame every %v seconds.", frameInterval)
	log.Printf("Aggregating new face counts over %v intervals.", aggregationInterval)

	// ------------------------------------------------
	// 6. Main loop: Capture, detect, display, and aggregate new faces
	// ------------------------------------------------
	for {
		if ok := cap.Read(&img); !ok {
			log.Println("Cannot read frame from camera. Exiting.")
			break
		}
		if img.Empty() {
			time.Sleep(frameInterval)
			continue
		}

		// Create a blob from the image for DNN input (resize to 300x300, subtract mean)
		blob := gocv.BlobFromImage(img, 1.0, image.Pt(300, 300),
			gocv.NewScalar(104.0, 177.0, 123.0, 0), false, false)
		net.SetInput(blob, "")
		detections := net.Forward("")
		blob.Close()

		// Reshape the detections to a 2D matrix where each row is one detection.
		totalDetections := int(detections.Total() / 7)
		detectionMat := detections.Reshape(1, totalDetections)

		// currentDetections holds all bounding boxes detected in the current frame.
		var currentDetections []image.Rectangle
		newFaces := 0

		// Loop over each detection
		for i := 0; i < totalDetections; i++ {
			confidence := detectionMat.GetFloatAt(i, 2)
			if confidence < float32(dnnConfidenceThreshold) {
				continue
			}

			// Compute bounding box coordinates
			left := int(detectionMat.GetFloatAt(i, 3) * float32(img.Cols()))
			top := int(detectionMat.GetFloatAt(i, 4) * float32(img.Rows()))
			right := int(detectionMat.GetFloatAt(i, 5) * float32(img.Cols()))
			bottom := int(detectionMat.GetFloatAt(i, 6) * float32(img.Rows()))
			rect := image.Rect(left, top, right, bottom)
			currentDetections = append(currentDetections, rect)

			// Check against previous detections to see if it's a duplicate.
			duplicate := false
			for _, prev := range previousDetections {
				if IoU(rect, prev) >= duplicateThreshold {
					duplicate = true
					break
				}
			}
			if !duplicate {
				newFaces++
			}

			// Draw the detection box and confidence label.
			gocv.Rectangle(&img, rect, color.RGBA{0, 255, 0, 0}, 2)
			label := fmt.Sprintf("%.2f", confidence)
			size := gocv.GetTextSize(label, gocv.FontHersheySimplex, 0.5, 1)
			pt := image.Pt(left, top-2)
			gocv.Rectangle(&img, image.Rect(pt.X, pt.Y-size.Y, pt.X+size.X, pt.Y), color.RGBA{0, 255, 0, 0}, -1)
			gocv.PutText(&img, label, pt, gocv.FontHersheySimplex, 0.5, color.RGBA{0, 0, 0, 0}, 1)
		}

		// Update aggregated count with only new faces from this frame.
		aggregatedCount += newFaces
		log.Printf("Detected %d new face(s) in this frame (total new faces: %d).", newFaces, aggregatedCount)

		// Update previous detections to current ones for next frame comparison.
		previousDetections = currentDetections

		// Display the frame with detections
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}

		// Check if the aggregation interval has passed
		if time.Since(startTime) >= aggregationInterval {
			log.Printf("Aggregation interval reached. Aggregated new face count = %d", aggregatedCount)
			err = sendDataToThingSpeak(thingspeakURL, thingspeakWriteKey, aggregatedCount)
			if err != nil {
				log.Println("Error sending data to ThingSpeak:", err)
			} else {
				log.Println("Data successfully sent to ThingSpeak.")
			}
			aggregatedCount = 0
			startTime = time.Now()
			// Optionally, clear previous detections here to reset the tracking.
			previousDetections = nil
		}

		time.Sleep(frameInterval)
	}
}

// getEnv returns the environment variable value if it exists,
// otherwise it returns the fallback value.
func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
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

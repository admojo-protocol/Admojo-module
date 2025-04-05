# AdNet Protocol - Smart Advertising with IoT + AI

**Flow**: ESP32-CAM → Stream Video → Detection using RNN and duplication avoidance using IoU → Aggregated # of viewers in 5 min durations → Send to ThingSpeak.

---

## 📸 ESP32-CAM Module: AdNetCamModule

A tiny AI Thinker ESP32-CAM module powers our eyes-on-the-ground (literally). It streams MJPEG video over Wi-Fi so we can detect who's watching those juicy ads in real-time.

### 🔧 Hardware Used
- **Board:** ESP32-CAM (AI Thinker variant with PSRAM)
- **Camera:** OV2640
- **Status LED Pin:** GPIO 33

### ⚙️ Features
- MJPEG video streaming over Wi-Fi (on http://192.168.1.43/)
- Static IP config for consistency
- Status LED blinks while connecting, stays off when ready
- Stream served as `multipart/x-mixed-replace;boundary=123456789000000000000987654321`
- Optimized for high-quality JPEG (UXGA or SVGA)

### 📝 Code Summary
See `AdNetCamModule.ino` for full source. Highlights:
- Initializes camera with dynamic config (PSRAM-aware)
- Sets up MJPEG stream endpoint `/`
- Streams images using ESP-IDF HTTP server
- Loop does nothing; all work is done in background

**To Access the Stream:**
```bash
http://192.168.1.43/  # Served MJPEG Stream
```

---

## 🤖 Backend: Go + OpenCV + AI = ✨

### 🔄 Flow
1. Go app connects to MJPEG stream from ESP32-CAM
2. Every 5 seconds, it grabs a frame
3. Runs OpenCV DNN face detection (SSD + ResNet)
4. Tracks new viewers (avoids counting duplicates using IoU filtering)
5. Aggregates counts for 5 minutes
6. Sends the count to [ThingSpeak](https://thingspeak.com/) via API

### 🧩 Features
- Uses GoCV (Go bindings for OpenCV)
- DNN-based detector (ResNet SSD) for better accuracy than Haar
- Deduplicates people using bounding-box IoU tracking
- Configuration via `.env` (stream URL, model paths, thresholds)
- Displays debug video with rectangles and confidence
- Uses ThingSpeak to report proof-of-view metrics

### 🛠⃣ How It Works
- On every frame, we detect faces
- We compare new bounding boxes to the previous frame
- If IoU < threshold, it’s a new person (so we count it!)
- Every 5 minutes, we send the total to ThingSpeak

### 👀 Want to See the AI in Action?
You’ll get a debug window with live rectangles drawn over faces along with confidence scores.

---

## 📍 The Result?
- Transparent, real-world Proof of Views
- Performance-based rewards for Ad Service Providers
- Smarter ads that actually engage people
- All running on open tech you control

---

## ✅ What's Next?
- Plug into smart contracts for dynamic ad payouts
- Integrate cryptography for verifiability on-chain

Stay tuned. This is just the beginning.

Made with ❤️ by Anon.
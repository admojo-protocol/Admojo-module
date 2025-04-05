#include "esp_camera.h"
#include <WiFi.h>

extern "C" {
  #include "esp_http_server.h"
}
// Include our helper libraries
#include "FirmwareHash.h"
#include "SecureSign.h"

// ===================
// Select camera model
// ===================
#define CAMERA_MODEL_AI_THINKER // Has PSRAM

#include "camera_pins.h"

// ===========================
// Enter your WiFi credentials
// ===========================
const char* ssid     = "FTTH_2.4ghz";
const char* password = "wefwegefgwe";

// ===========================
// Status LED pin configuration
// ===========================

#define LED_STATUS_PIN 33

// ---------------------------
// Wi-Fi Static IP parameters
// ---------------------------
IPAddress local_IP(192, 168, 1, 43);
IPAddress gateway(192, 168, 1, 1);
IPAddress subnet(255, 255, 255, 0);

IPAddress primaryDNS(8, 8, 8, 8);
IPAddress secondaryDNS(8, 8, 4, 4);

// Forward declarations
static esp_err_t stream_handler(httpd_req_t *req);
void startCameraServer();
static void blinkWifiConnecting(void);

httpd_handle_t stream_httpd = NULL;

// Global strings to hold the firmware hash and signature in hex format.
String gFirmwareHashHex = "";
String gSignatureHex = "";

// -----------------------------------------------------------
// Blink LED while connecting to Wi-Fi
// -----------------------------------------------------------
static void blinkWifiConnecting(void)
{
  static unsigned long lastBlink = 0;
  static bool ledState = false;
  if (millis() - lastBlink >= 300) {
    lastBlink = millis();
    ledState = !ledState;
    digitalWrite(LED_STATUS_PIN, ledState ? HIGH : LOW);
  }
}

// -----------------------------------------------------------
// Helper: Convert byte array to hex string
// -----------------------------------------------------------
String bytesToHexString(const uint8_t* data, size_t len) {
  String hexStr = "";
  for (size_t i = 0; i < len; i++) {
    if (data[i] < 16) hexStr += "0";
    hexStr += String(data[i], HEX);
  }
  hexStr.toUpperCase();
  return hexStr;
}

// -----------------------------------------------------------
// Setup
// -----------------------------------------------------------
void setup()
{
  Serial.begin(115200);
  Serial.setDebugOutput(true);
  Serial.println();
  pinMode(LED_STATUS_PIN, OUTPUT);
  digitalWrite(LED_STATUS_PIN, LOW);

  // Configure static IP.
  if (!WiFi.config(local_IP, gateway, subnet, primaryDNS, secondaryDNS)) {
    Serial.println("STA Failed to configure");
    digitalWrite(LED_STATUS_PIN, HIGH); // steady LED = error
    while (true) { delay(100); }
  }

  WiFi.begin(ssid, password);
  WiFi.setSleep(false);
  Serial.println("WiFi connecting...");

  while (WiFi.status() != WL_CONNECTED) {
    blinkWifiConnecting();
    delay(50);
  }

  digitalWrite(LED_STATUS_PIN, LOW);
  Serial.printf("\nWiFi connected. IP address: %s\n", WiFi.localIP().toString().c_str());

  // -----------------------------------------------------------
  // Camera configuration
  // -----------------------------------------------------------
  camera_config_t config;
  config.ledc_channel = LEDC_CHANNEL_0;
  config.ledc_timer   = LEDC_TIMER_0;
  config.pin_d0       = Y2_GPIO_NUM;
  config.pin_d1       = Y3_GPIO_NUM;
  config.pin_d2       = Y4_GPIO_NUM;
  config.pin_d3       = Y5_GPIO_NUM;
  config.pin_d4       = Y6_GPIO_NUM;
  config.pin_d5       = Y7_GPIO_NUM;
  config.pin_d6       = Y8_GPIO_NUM;
  config.pin_d7       = Y9_GPIO_NUM;
  config.pin_xclk     = XCLK_GPIO_NUM;
  config.pin_pclk     = PCLK_GPIO_NUM;
  config.pin_vsync    = VSYNC_GPIO_NUM;
  config.pin_href     = HREF_GPIO_NUM;
  config.pin_sccb_sda = SIOD_GPIO_NUM;
  config.pin_sccb_scl = SIOC_GPIO_NUM;
  config.pin_pwdn     = PWDN_GPIO_NUM;
  config.pin_reset    = RESET_GPIO_NUM;
  config.xclk_freq_hz = 20000000;
  config.frame_size    = FRAMESIZE_UXGA;   // 1600x1200 (high quality)
  config.pixel_format  = PIXFORMAT_JPEG;     // for streaming
  config.grab_mode     = CAMERA_GRAB_WHEN_EMPTY;
  config.fb_location   = CAMERA_FB_IN_PSRAM;
  config.jpeg_quality  = 12;
  config.fb_count      = 1;

  if (psramFound()) {
    config.jpeg_quality = 10;
    config.fb_count = 2;
    config.grab_mode = CAMERA_GRAB_LATEST;
  } else {
    config.frame_size  = FRAMESIZE_SVGA;  // 800x600 if no PSRAM
    config.fb_location = CAMERA_FB_IN_DRAM;
  }

  esp_err_t err = esp_camera_init(&config);
  if (err != ESP_OK) {
    Serial.printf("Camera init failed with error 0x%x\n", err);
    digitalWrite(LED_STATUS_PIN, HIGH);
    while (true) { delay(100); }
  }

  sensor_t *s = esp_camera_sensor_get();
  if (s->id.PID == OV3660_PID) {
    s->set_vflip(s, 1);
    s->set_brightness(s, 1);
    s->set_saturation(s, -2);
  }

  // Start HTTPS camera server (ensure you configure TLS certs in httpd_config_t if needed)
  startCameraServer();
  Serial.println("Camera Ready! Visit: https://" + WiFi.localIP().toString());

  // -----------------------------------------------------------
  // Compute firmware hash and sign it.
  // -----------------------------------------------------------
  uint8_t fwHash[32];
  if (computeFirmwareHash(fwHash)) {
    gFirmwareHashHex = bytesToHexString(fwHash, sizeof(fwHash));
    Serial.print("Firmware Hash: ");
    Serial.println(gFirmwareHashHex);

    // Sign the firmware hash using our secure key.
    uint8_t signature[65]; // 65 bytes: 32 (r) + 32 (s) + 1 (v)
    size_t sigLen = 0;
    if (signData(fwHash, sizeof(fwHash), signature, &sigLen)) {
      gSignatureHex = bytesToHexString(signature, sigLen);
      Serial.println("Firmware hash signed successfully!");
      Serial.print("Signature (hex): ");
      Serial.println(gSignatureHex);
    } else {
      Serial.println("Failed to sign the firmware hash.");
    }
  } else {
    Serial.println("Failed to compute firmware hash!");
  }
}

void loop()
{
  // Server runs in the background.

// -----------------------------------------------------------
// MJPEG Stream Handler (with extra headers)
// -----------------------------------------------------------
static esp_err_t stream_handler(httpd_req_t *req)
{
  // Add custom headers for on-chain verification.
  httpd_resp_set_hdr(req, "X-Firmware-Hash", gFirmwareHashHex.c_str());
  httpd_resp_set_hdr(req, "X-Signature", gSignatureHex.c_str());

  // Set content type to MJPEG.
  static const char* CONTENT_TYPE = "multipart/x-mixed-replace;boundary=123456789000000000000987654321";
  static const char* BOUNDARY     = "\r\n--123456789000000000000987654321\r\n";
  static const char* PART_FMT     = "Content-Type: image/jpeg\r\nContent-Length: %u\r\nX-Timestamp: %d.%06d\r\n\r\n";

  esp_err_t res = httpd_resp_set_type(req, CONTENT_TYPE);
  if (res != ESP_OK) return res;

  httpd_resp_set_hdr(req, "Access-Control-Allow-Origin", "*");

  while (true) {
    camera_fb_t *fb = esp_camera_fb_get();
    if (!fb) {
      Serial.println("Camera capture failed");
      res = ESP_FAIL;
    } else {
      uint8_t* jpg_buf = NULL;
      size_t jpg_buf_len = 0;
      if (fb->format != PIXFORMAT_JPEG) {
        bool converted = frame2jpg(fb, 80, &jpg_buf, &jpg_buf_len);
        esp_camera_fb_return(fb);
        fb = NULL;
        if (!converted) {
          Serial.println("JPEG compression failed");
          res = ESP_FAIL;
        }
      } else {
        jpg_buf_len = fb->len;
        jpg_buf     = fb->buf;
      }

      if (res == ESP_OK) {
        res = httpd_resp_send_chunk(req, BOUNDARY, strlen(BOUNDARY));
      }
      if (res == ESP_OK) {
        char part_buf[128];
        size_t hlen = snprintf(part_buf, sizeof(part_buf), PART_FMT,
                               jpg_buf_len,
                               fb ? fb->timestamp.tv_sec : 0,
                               fb ? fb->timestamp.tv_usec : 0);
        res = httpd_resp_send_chunk(req, part_buf, hlen);
      }
      if (res == ESP_OK) {
        res = httpd_resp_send_chunk(req, (const char*)jpg_buf, jpg_buf_len);
      }

      if (fb) {
        esp_camera_fb_return(fb);
      } else if (jpg_buf) {
        free(jpg_buf);
      }
    }

    if (res != ESP_OK) {
      break;
    }
  }
  return res;
}

// -----------------------------------------------------------
// Start Camera Server
// -----------------------------------------------------------
void startCameraServer()
{
  httpd_config_t config = HTTPD_DEFAULT_CONFIG();
  // For HTTPS support, we must configure TLS certificates in config.
  if (httpd_start(&stream_httpd, &config) == ESP_OK) {
    httpd_uri_t stream_uri = {
      .uri       = "/",
      .method    = HTTP_GET,
      .handler   = stream_handler,
      .user_ctx  = NULL
    };
    httpd_register_uri_handler(stream_httpd, &stream_uri);
    Serial.println("HTTP Server started on port " + String(config.server_port));
  } else {
    Serial.println("Failed to start HTTP server");
    // Turn LED on to indicate server start error
    digitalWrite(LED_STATUS_PIN, HIGH);
  }
}

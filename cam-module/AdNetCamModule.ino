#include "esp_camera.h"
#include <WiFi.h>

extern "C" {
  #include "esp_http_server.h"
}
// ===================
// Select camera model
// ===================
#define CAMERA_MODEL_AI_THINKER // Has PSRAM

#include "camera_pins.h"

// ===========================
// Enter your WiFi credentials
// ===========================
const char* ssid     = "FTTH_2.4ghz";
const char* password = "qefe2fqe";

// ===========================
// Status LED pin configuration
// ===========================
// Change this to a valid pin for a status LED on your board.
// If you do not have a spare LED, you can disable these lines or
// re-map them. Many ESP32-CAM boards do not have a dedicated status LED.
#define LED_STATUS_PIN 33

// ---------------------------
// Wi-Fi Static IP parameters
// ---------------------------
IPAddress local_IP(192, 168, 1, 43);
IPAddress gateway(192, 168, 1, 1);
IPAddress subnet(255, 255, 255, 0);
IPAddress primaryDNS(8, 8, 8, 8);    // optional
IPAddress secondaryDNS(8, 8, 4, 4);  // optional

// Forward declarations
static esp_err_t stream_handler(httpd_req_t *req);
void startCameraServer();
static void blinkWifiConnecting(void);

httpd_handle_t stream_httpd = NULL;

// -----------------------------------------------------------
// Blink the LED rapidly while trying to connect to Wi-Fi
// -----------------------------------------------------------
static void blinkWifiConnecting(void)
{
  static unsigned long lastBlink = 0;
  static bool ledState = false;
  if (millis() - lastBlink >= 300) { // adjust blink speed if desired
    lastBlink = millis();
    ledState = !ledState;
    digitalWrite(LED_STATUS_PIN, ledState ? HIGH : LOW);
  }
}

// -----------------------------------------------------------
// Setup function
// -----------------------------------------------------------
void setup()
{
  Serial.begin(115200);
  Serial.setDebugOutput(true);
  Serial.println();

  // Set up the LED (for status) and turn it off initially
  pinMode(LED_STATUS_PIN, OUTPUT);
  digitalWrite(LED_STATUS_PIN, LOW);

  // Static IP configuration
  if (!WiFi.config(local_IP, gateway, subnet, primaryDNS, secondaryDNS)) {
    Serial.println("STA Failed to configure");
    // Turn LED steady on to indicate error
    digitalWrite(LED_STATUS_PIN, HIGH);
    while (true) {
      delay(100);
    }
  }

  WiFi.begin(ssid, password);
  WiFi.setSleep(false);

  Serial.println("WiFi connecting...");

  // Blink LED while connecting
  while (WiFi.status() != WL_CONNECTED) {
    blinkWifiConnecting();
    delay(50);
  }

  // Once connected, turn the LED off (indicating normal operation)
  digitalWrite(LED_STATUS_PIN, LOW);

  Serial.println("");
  Serial.println("WiFi connected. IP address: ");
  Serial.println(WiFi.localIP());

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

  // High-quality capture settings (UXGA + JPEG)
  config.frame_size    = FRAMESIZE_UXGA;   // 1600x1200 (UXGA)
  config.pixel_format  = PIXFORMAT_JPEG;  // for streaming
  config.grab_mode     = CAMERA_GRAB_WHEN_EMPTY;
  config.fb_location   = CAMERA_FB_IN_PSRAM;
  config.jpeg_quality  = 12;
  config.fb_count      = 1;

  // If PSRAM is found, improve performance/quality
  if (psramFound()) {
    config.jpeg_quality = 10;             // better quality
    config.fb_count = 2;                  // double-buffer
    config.grab_mode = CAMERA_GRAB_LATEST;
  } else {
    // If no PSRAM, reduce resolution
    config.frame_size  = FRAMESIZE_SVGA;  // 800x600
    config.fb_location = CAMERA_FB_IN_DRAM;
  }

  // -----------------------------------------------------------
  // Initialize camera
  // -----------------------------------------------------------
  esp_err_t err = esp_camera_init(&config);
  if (err != ESP_OK) {
    Serial.printf("Camera init failed with error 0x%x", err);
    // Turn LED steady on to indicate camera error
    digitalWrite(LED_STATUS_PIN, HIGH);
    while (true) {
      delay(100);
    }
  }
  
  // Optional sensor adjustments
  sensor_t *s = esp_camera_sensor_get();
  if (s->id.PID == OV3660_PID) {
    s->set_vflip(s, 1);        // flip it back
    s->set_brightness(s, 1);   // up the brightness
    s->set_saturation(s, -2);  // lower the saturation
  }

  // DO NOT drop down to QVGA—keep it at UXGA or SVGA as configured above
  // e.g. remove the typical line:
  // s->set_framesize(s, FRAMESIZE_QVGA); // <--- REMOVED

  // Start the camera server
  startCameraServer();

  Serial.print("Camera Ready! Use 'http://");
  Serial.print(WiFi.localIP());
  Serial.println("' to connect (MJPEG stream).");
}

void loop()
{
  // The server runs in the background. Nothing special needed here.
  delay(1000);
}

// -----------------------------------------------------------
// MJPEG Stream Handler (only endpoint)
// -----------------------------------------------------------
static esp_err_t stream_handler(httpd_req_t *req)
{
  // Provide MJPEG "multipart/x-mixed-replace" content
  static const char*  _STREAM_CONTENT_TYPE = "multipart/x-mixed-replace;boundary=123456789000000000000987654321";
  static const char*  _STREAM_BOUNDARY     = "\r\n--123456789000000000000987654321\r\n";
  static const char*  _STREAM_PART        = "Content-Type: image/jpeg\r\nContent-Length: %u\r\nX-Timestamp: %d.%06d\r\n\r\n";

  camera_fb_t * fb = NULL;
  esp_err_t res = ESP_OK;
  uint8_t * _jpg_buf = NULL;
  size_t _jpg_buf_len = 0;
  char part_buf[128];

  // Set appropriate content type
  res = httpd_resp_set_type(req, _STREAM_CONTENT_TYPE);
  if (res != ESP_OK) {
    return res;
  }

  // We can set a custom header if desired
  httpd_resp_set_hdr(req, "Access-Control-Allow-Origin", "*");

  // Streaming loop
  while (true) {
    fb = esp_camera_fb_get();
    if (!fb) {
      Serial.println("Camera capture failed");
      res = ESP_FAIL;
    } else {
      if (fb->format != PIXFORMAT_JPEG) {
        bool jpeg_converted = frame2jpg(fb, 80, &_jpg_buf, &_jpg_buf_len);
        esp_camera_fb_return(fb);
        fb = NULL;
        if (!jpeg_converted) {
          Serial.println("JPEG compression failed");
          res = ESP_FAIL;
        }
      } else {
        _jpg_buf_len = fb->len;
        _jpg_buf = fb->buf;
      }
    }

    if (res == ESP_OK) {
      // Send boundary
      res = httpd_resp_send_chunk(req, _STREAM_BOUNDARY, strlen(_STREAM_BOUNDARY));
    }
    if (res == ESP_OK) {
      // Prepare the MJPEG part header
      size_t hlen = snprintf(part_buf, 128, _STREAM_PART,
                             _jpg_buf_len,
                             fb ? fb->timestamp.tv_sec : 0,
                             fb ? fb->timestamp.tv_usec : 0);
      // Send header
      res = httpd_resp_send_chunk(req, part_buf, hlen);
    }
    if (res == ESP_OK) {
      // Send frame data
      res = httpd_resp_send_chunk(req, (const char*)_jpg_buf, _jpg_buf_len);
    }

    // Return the frame buffer back to be reused
    if (fb) {
      esp_camera_fb_return(fb);
      fb = NULL;
      _jpg_buf = NULL;
    } else if (_jpg_buf) {
      free(_jpg_buf);
      _jpg_buf = NULL;
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
  // Optionally, adjust server port if needed:
  // config.server_port = 80; // default

  // Start the web server
  if (httpd_start(&stream_httpd, &config) == ESP_OK) {
    // Register only one URI: the streaming endpoint at "/"
    httpd_uri_t stream_uri = {
      .uri       = "/",             // <--- single endpoint
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

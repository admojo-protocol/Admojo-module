#include <SPI.h>
// #include <WiFi.h>         // <-- Uncomment when you're ready to use Wi-Fi
// #include <HTTPClient.h>   // <-- Uncomment when you're ready to fetch URLs
#include <Adafruit_PN532.h>

// =====================================================================
// Pin/Constructor Setup
// =====================================================================
#define PN532_SCK    18
#define PN532_MISO   19
#define PN532_MOSI   23
#define PN532_SS     5

// Hardware Reset Pin for PN532 (RSTPD_N pin)
#define PN532_RESET  15   // Modify this to your actual GPIO pin for RSTPD_N

// Status LED for visual feedback (if available)
#define STATUS_LED   2

Adafruit_PN532 nfc(PN532_SCK, PN532_MISO, PN532_MOSI, PN532_SS);

// =====================================================================
// Keys for Mifare Classic
// =====================================================================
uint8_t keyA_sector1[6] = { 0xD3, 0xF7, 0xD3, 0xF7, 0xD3, 0xF7 };  // For sector 1..15
uint8_t keyB_all[6]     = { 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF };  // Common Key B

// =====================================================================
// Wi-Fi Credentials (Commented out for now)
// =====================================================================
/*
const char* ssid     = "YOUR_SSID";
const char* password = "YOUR_PASSWORD";
*/

// =====================================================================
// Dummy Array of Test URLs to pick from at random (for demonstration)
// =====================================================================
const char* testUrls[] = {
  "youtube.com",
  "aviral.software",
  "google.com",
  "some-other-url.xyz"
};
const uint8_t NUM_TEST_URLS = sizeof(testUrls) / sizeof(testUrls[0]);

// =====================================================================
// Forward Declarations
// =====================================================================
void enablePN532();
void disablePN532();
bool writeUrlToTag(const char* newUrl);
void goToDeepSleep(uint32_t seconds);

// =====================================================================
// Setup
// =====================================================================
void setup() {
  Serial.begin(115200);
  delay(1000); // Give time for serial to connect
  
  // Setup LED if available
  pinMode(STATUS_LED, OUTPUT);
  digitalWrite(STATUS_LED, HIGH);

  // Setup the PN532 RESET (RSTPD_N) pin
  pinMode(PN532_RESET, OUTPUT);
  
  Serial.println("\n\n=== AdNet NFC Writer (Hard-Power-Down with Random URLs) ===");

  // (Optional) Seed the random number generator
  // If you have a floating analog pin, you can use it. Otherwise, you can use micros().
  // e.g. randomSeed(analogRead(A0)); // or:
  randomSeed(micros());
  
  // (Optional) Connect to Wi-Fi and fetch new URL from server (commented out)
  /*
  Serial.println("Connecting to Wi-Fi...");
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("\nWiFi connected!");
  
  // Example of HTTP GET request (commented out)
  HTTPClient http;
  http.begin("https://admojo.xyz/advertiser/[adid]");  // Replace with your actual endpoint
  int httpCode = http.GET();
  if(httpCode == 200) {
    String payload = http.getString();
    Serial.println("Server response: " + payload);
    // parse JSON or manipulate 'payload' here to get the URL
  } else {
    Serial.print("HTTP GET failed, error: ");
    Serial.println(httpCode);
  }
  http.end();
  */

  // For demonstration, pick a random URL from our test array
  uint8_t randomIndex = random(0, NUM_TEST_URLS);
  const char* newUrl = testUrls[randomIndex];
  
  // Enable PN532 for tag writing
  enablePN532();

  // Attempt to write to NFC tag
  Serial.print("Randomly selected URL: ");
  Serial.println(newUrl);
  Serial.println("Looking for tag to write URL...");
  
  bool success = writeUrlToTag(newUrl);
  if (success) {
    Serial.println("Successfully wrote new URL to the NFC tag!");

    // Success indicator - blink LED rapidly 3 times
    for (int i = 0; i < 3; i++) {
      digitalWrite(STATUS_LED, LOW);
      delay(100);
      digitalWrite(STATUS_LED, HIGH);
      delay(100);
    }
  } else {
    Serial.println("Failed to write URL to the NFC tag...");

    // Failure indicator - blink LED slowly 2 times
    for (int i = 0; i < 2; i++) {
      digitalWrite(STATUS_LED, LOW);
      delay(300);
      digitalWrite(STATUS_LED, HIGH);
      delay(300);
    }
  }

  // Disable PN532 to allow phones to read the tag
  Serial.println("Disabling PN532 to allow phones to read the tag...");
  disablePN532();

  // Turn off status LED after finishing
  digitalWrite(STATUS_LED, LOW);

  // Go to deep sleep for 5 minutes (testing) or 1 hour in production
  goToDeepSleep(5 * 60);        // 5 minutes for demo
  // goToDeepSleep(60 * 60);    // 1 hour for production
}

// =====================================================================
// Main loop does nothing; we use deep sleep
// =====================================================================
void loop() {
  // This will not run because we enter deep sleep in setup(),
  // and upon wake, setup() runs fresh again.
}

// =====================================================================
// Enable PN532 using hardware reset
// =====================================================================
void enablePN532() {
  Serial.println("Enabling PN532...");

  // According to PN532 datasheet or Adafruit recommendations:
  // Bring the RSTPD_N pin HIGH to turn PN532 on.
  digitalWrite(PN532_RESET, HIGH);
  delay(100);

  // Initialize PN532
  nfc.begin();

  // Check if PN532 is responding
  uint32_t versiondata = nfc.getFirmwareVersion();
  if (!versiondata) {
    Serial.println("Did not find PN532 board! Retrying...");

    // Try a hardware reset again
    digitalWrite(PN532_RESET, LOW);
    delay(100);
    digitalWrite(PN532_RESET, HIGH);
    delay(100);

    nfc.begin();
    versiondata = nfc.getFirmwareVersion();

    if (!versiondata) {
      Serial.println("Still cannot find PN532. Check wiring.");
      // Error indicator: blink LED forever
      while (1) {
        digitalWrite(STATUS_LED, !digitalRead(STATUS_LED));
        delay(300);
      }
    }
  }

  Serial.print("Found PN532 with firmware version: ");
  Serial.print((versiondata >> 16) & 0xFF, DEC);
  Serial.print(".");
  Serial.println((versiondata >> 8) & 0xFF, DEC);

  // Configure PN532 for Mifare Classic
  nfc.setPassiveActivationRetries(0xFF);
  nfc.SAMConfig();

  Serial.println("PN532 enabled and ready!");
}

// =====================================================================
// Disable PN532 using hardware reset (Hard Power Down)
// =====================================================================
void disablePN532() {
  Serial.println("Disabling PN532 with hardware reset (pull RSTPD_N LOW)...");
  
  // Pull RSTPD_N pin LOW to disable (power down) PN532
  digitalWrite(PN532_RESET, LOW);
  delay(50);

  Serial.println("PN532 disabled. Phones should now be able to read the tag easily.");
}

// =====================================================================
// Write a new URL to the tag in blocks 4 and 5 (Sector 1)
// =====================================================================
bool writeUrlToTag(const char* newUrl) {
  // Wait up to 5 seconds for a card to appear
  uint8_t uid[7]   = { 0 };
  uint8_t uidLen   = 0;
  bool cardFound   = false;

  long startTime = millis();
  while ((millis() - startTime) < 5000) {
    if (nfc.readPassiveTargetID(PN532_MIFARE_ISO14443A, uid, &uidLen)) {
      cardFound = true;
      break;
    }
    delay(100);
    
    // Flash the LED so user sees "searching"
    if ((millis() - startTime) % 500 < 250) {
      digitalWrite(STATUS_LED, HIGH);
    } else {
      digitalWrite(STATUS_LED, LOW);
    }
  }

  // Turn LED back on after searching
  digitalWrite(STATUS_LED, HIGH);

  if (!cardFound) {
    Serial.println("No Mifare Classic card found within 5 seconds.");
    return false;
  }

  Serial.print("Card found! UID: ");
  for (uint8_t i = 0; i < uidLen; i++) {
    Serial.print(uid[i], HEX);
    Serial.print(" ");
  }
  Serial.println();

  // Construct the NDEF TLV
  // For "https://www." prefix => 0x02
  // e.g., newUrl = "youtube.com" => phone sees "https://www.youtube.com"
  uint8_t urlLen = strlen(newUrl);
  uint8_t payloadLen     = 1 + urlLen;       // 1 for the prefix byte
  uint8_t ndefRecordLen  = 5 + urlLen;       // 0xD1, 0x01, payloadLen, 0x55, 0x02, plus the url
  uint8_t totalLen       = ndefRecordLen + 3;// plus 0x03 (TLV Tag), length byte, and 0xFE (terminator)

  // Prepare a buffer for blocks 4 and 5 (32 bytes max)
  uint8_t ndefBuf[32];
  memset(ndefBuf, 0, 32);

  // Fill the TLV
  int idx = 0;
  ndefBuf[idx++] = 0x03;                // NDEF Tag
  ndefBuf[idx++] = ndefRecordLen;       // length of the NDEF record
  // NDEF record
  ndefBuf[idx++] = 0xD1;                // MB=1, ME=1, SR=1, TNF=1 (well-known)
  ndefBuf[idx++] = 0x01;                // type length = 1 ('U')
  ndefBuf[idx++] = payloadLen;          // payload length
  ndefBuf[idx++] = 0x55;                // 'U'
  ndefBuf[idx++] = 0x02;                // "https://www." prefix
  // Copy the url
  for (uint8_t i = 0; i < urlLen; i++) {
    ndefBuf[idx++] = newUrl[i];
  }
  // Terminator
  ndefBuf[idx++] = 0xFE;

  // Split into two blocks (16 bytes each)
  uint8_t block4[16], block5[16];
  memset(block4, 0, 16);
  memset(block5, 0, 16);

  uint8_t copyLenBlock4 = (totalLen > 16) ? 16 : totalLen;
  memcpy(block4, ndefBuf, copyLenBlock4);
  if (totalLen > 16) {
    uint8_t remainder = totalLen - 16;
    if (remainder > 16) remainder = 16;  // if needed, you'd go to block 6, etc.
    memcpy(block5, ndefBuf + 16, remainder);
  }

  // Attempt up to 3 writes in case of authentication or write failures
  bool success = false;
  for (int attempt = 0; attempt < 3 && !success; attempt++) {
    // Authenticate and write block 4
    if (!nfc.mifareclassic_AuthenticateBlock(uid, uidLen, 4, 0, keyA_sector1)) {
      Serial.print("Auth failed for block 4 (attempt ");
      Serial.print(attempt + 1);
      Serial.println(")");
      delay(200);
      continue;
    }
    if (!nfc.mifareclassic_WriteDataBlock(4, block4)) {
      Serial.print("Write failed for block 4 (attempt ");
      Serial.print(attempt + 1);
      Serial.println(")");
      delay(200);
      continue;
    }

    // Authenticate and write block 5
    if (!nfc.mifareclassic_AuthenticateBlock(uid, uidLen, 5, 0, keyA_sector1)) {
      Serial.print("Auth failed for block 5 (attempt ");
      Serial.print(attempt + 1);
      Serial.println(")");
      delay(200);
      continue;
    }
    if (!nfc.mifareclassic_WriteDataBlock(5, block5)) {
      Serial.print("Write failed for block 5 (attempt ");
      Serial.print(attempt + 1);
      Serial.println(")");
      delay(200);
      continue;
    }

    // Verification (optional but recommended)
    uint8_t readBlock[16];
    bool verifyOk = true;

    // Verify block 4
    if (!nfc.mifareclassic_AuthenticateBlock(uid, uidLen, 4, 0, keyA_sector1) ||
        !nfc.mifareclassic_ReadDataBlock(4, readBlock)) {
      Serial.println("Verification failed: couldn't read block 4");
      verifyOk = false;
    } else {
      for (int i = 0; i < 16; i++) {
        if (readBlock[i] != block4[i]) {
          Serial.println("Verification failed: block 4 data mismatch");
          verifyOk = false;
          break;
        }
      }
    }

    // Verify block 5
    if (verifyOk) {
      if (!nfc.mifareclassic_AuthenticateBlock(uid, uidLen, 5, 0, keyA_sector1) ||
          !nfc.mifareclassic_ReadDataBlock(5, readBlock)) {
        Serial.println("Verification failed: couldn't read block 5");
        verifyOk = false;
      } else {
        for (int i = 0; i < 16; i++) {
          if (readBlock[i] != block5[i]) {
            Serial.println("Verification failed: block 5 data mismatch");
            verifyOk = false;
            break;
          }
        }
      }
    }

    if (verifyOk) {
      success = true;
      Serial.print("NDEF URL updated and verified: https://www.");
      Serial.println(newUrl);
    } else {
      Serial.println("Verification failed, retrying...");
      delay(500);
    }
  }

  return success;
}

// =====================================================================
// Put ESP32 into deep sleep
// =====================================================================
void goToDeepSleep(uint32_t seconds) {
  Serial.printf("Going to deep sleep for %u seconds...\n", seconds);
  Serial.flush();

  esp_sleep_enable_timer_wakeup((uint64_t)seconds * 1000000ULL);
  esp_deep_sleep_start();
}

# 🌐 ADMOJO Protocol  
### *Smarter Ads. Cryptographically Verified Engagement.*

Welcome to **ADMOJO Protocol: Module Subrepo**

Say goodbye to ad fraud. With cryptographic **Proof of Views** (via ESP32-CAM) and **Proof of Taps** (via NFC interactions), ADMOJO ensures every interaction is authentic and verifiable.

---

## 🚀 Modules Breakdown

### 1. 📸 `admojo_web_server` — *Cryptographic Proof of Views Engine*

**Goal:** Authenticate real-time visual engagement using blockchain cryptography.

- **ESP32-CAM** securely streams MJPEG video
- Firmware authenticity ensured via SHA256 hashes signed using ECDSA (secp256k1)
- Golang backend (GoCV + OpenCV with SSD/ResNet DNN) detects human faces securely
- Aggregates cryptographically signed viewer counts, submitted securely on-chain

**Highlights:**
- SHA256 hash integrity verification
- IoU filtering prevents duplicate counts
- Real-time cryptographic verification of video feeds

> 🔍 Ads aren’t just viewed; they're cryptographically proven.

---

### 2. 📲 `admojo_nfc_module` — *Cryptographic Proof of Taps Module*

**Goal:** Ensure user taps are genuine and verifiable via blockchain.

- **ESP32 WROOM + PN532 NFC** securely fetches dynamic campaign URLs
- NFC tags written with NDEF format, secured using ECDSA-signed SHA256 firmware hashes
- Deep sleep cycles protect against interference and fraud

**Cryptographic Security:**
- Secure signing of URLs (ECDSA)
- NFC tags locked with unique cryptographic sector keys

> ✨ Passive views show interest; cryptographically secured taps prove engagement.

---

### 3. ⚙️ `admojo_compute_node` — *Golang Blockchain Oracle & Verifier*

**Goal:** Provide cryptographically verifiable data directly on Ethereum.

- Authenticates firmware hashes and signatures via Ethereum-compatible ECDSA (secp256k1)
- Real-time verification of SHA256 hashes ensures data integrity
- Aggregates viewer metrics, pushing cryptographically verified data to Ethereum smart contracts

**Core Features:**
- Robust ECDSA signature verification
- Secure integration with Ethereum blockchain
- Immutable on-chain storage of verified metrics

> 🔗 Securely bridging off-chain IoT interactions to Ethereum smart contracts.

---

## 🔗 On-Chain Cryptographic Logic

- **ADC Token:** Ethereum-based token staked and distributed via cryptographically secured engagement metrics
- **Smart Contracts:**
  - Verify ECDSA signatures and SHA256 firmware hashes in real-time
  - Automate token allocation based on cryptographically proven metrics
  - Transparent, immutable, and auditable Ethereum blockchain transactions

> All actions cryptographically **verifiable**, **immutable**, and **auditable**.

---

## 💡 Why ADMOJO?

- 🔒 **Cryptographic Verification:** Ensures authenticity and integrity
- 📡 **Blockchain IoT Integration:** Real-world secured analytics
- 🌐 **Ethereum Transparency:** Complete transparency in ad spend
- 💸 **Performance-Proven:** Pay exclusively for cryptographically verified engagement

---

## 📅 MVP Milestones

- ✅ **Subtask 1:** ESP32-CAM cryptographic views verification
- ✅ **Subtask 2:** NFC cryptographic security & dynamic campaign URLs
- ✅ **Subtask 3:** Golang backend cryptographic verification & Ethereum integration

---

## 🛠 Tech Stack

- **Hardware:** ESP32 WROOM/CAM, PN532 NFC Module
- **Software:** Arduino IDE, GoCV, OpenCV, Golang
- **Blockchain & Cryptography:** Ethereum, Solidity, Foundry, Metal API, ECDSA (secp256k1), SHA256

---

## 👾 Join the ADMOJO Movement

ADMOJO Protocol represents the next-generation Ethereum-powered, cryptographically secured advertising platform.

Join our community of blockchain developers, IoT enthusiasts, and cryptography innovators.

> **ADMOJO Protocol** — where interactions aren't merely recorded; they're cryptographically **proven**.

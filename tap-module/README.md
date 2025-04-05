# 🧠 AdMojo NFC Writer  
### *Power-efficient. Tamper-proof. Built for engagement.*

Welcome to the Admojo Protocol’s deep-sleep-powered NFC writer — the IoT module that doesn’t just sit there… it thinks, sleeps, writes, and resets the future of physical ad interactions.

## 🚀 What It Does

This device is the silent force behind **Proof of Taps** in the Admojo Protocol ecosystem — bridging the physical world of billboards and banners with the smart, accountable, on-chain world of ad monetization.

Every hour, it wakes up from deep sleep, securely updates an NFC tag with a **campaign-specific URL**, and vanishes into low-power hibernation, leaving behind a touchpoint that users can interact with using just their phones.

## ✨ Key Features

### 🎯 Tap-Triggered Engagement
Users tap the NFC tag placed on the ad display to instantly access live bounties, giveaways, or campaigns — each tap counts as a verified interaction and contributes to **Proof of Taps**, Admojo’s active participation metric.

### 🛡️ Secure By Design
The NFC tag is locked behind **cryptographically enforced access keys**, preventing unauthorized overwrites. Only Admojo-signed writers with validated keys can update tag content, ensuring **tamper-proof operations** in the wild.

### ⚡ Ultra Low Power, Ultra Smart
- Powered by an **ESP32 WROOM** module with built-in Wi-Fi
- Cycles through **deep sleep** for most of the hour
- Wakes only to write updates via **SPI-connected PN532**
- Uses **hard power-down** logic to shut off the PN532 module completely, eliminating RF field interference during idle time

### 🔁 Dynamic Ad Sync
- Fetches new campaign URLs via a secure GET request
- Writes NDEF-formatted links to NFC tags
- Integrates with on-chain reward systems and ThingSpeak-based oracle sync

### 🧠 On-chain readiness Ready
Each tap is captured as part of a **web-based analytics and reward pipeline** — feeding directly into Admojo’s smart contracts to:
- Reallocate ADC token rewards based on tap volume
- Optimize campaign ROI in real time
- Promote high-engagement ASP locations

### 📡 Seamless Web3 Integration
Works hand-in-hand with:
- **Admojo Campaign Backend**
- **Self Protocol** for secure, Sybil-resistant user verification and demographic analytics (age range, country, etc.) without compromising privacy through ZKProof ([Learn more](https://docs.self.xyz/))

### 🔐 Encrypted NFC Access Control
- Tag sectors are locked down with **access keys only known to the Admojo writer firmware**
- Unauthorized attempts to read or write are ignored or blocked
- Ensures that every update comes from a **trusted and signed source**

## 🧩 Use Case Snapshot

Imagine this:

> A user walks past a high-traffic LED ad board. Sees a message:  
> **"Tap here for 10 ADC!"**  
> They tap.  
> Their phone opens a URL embedded in the NFC tag — dynamically updated by the Admojo NFC Writer.  
> The tap is verified, logged, and counted.  
> Advertisers get accurate engagement metrics. ASPs earn rewards. Users securely verify themselves through Self Protocol and receive targeted offers.  
> Everyone wins.

## 🔮 Roadmap Enhancements

- Fully off-chain encrypted URL sync using Metal links.
- MQTT-injected push updates for real-time ad transitions (Nodit might help 👀)
- NFC + camera + Wifi (directional) fusion node for Proof-of-Presence triangulation  

---

## 🧬 The Bigger Picture

This isn’t a toy NFC writer.  
This is **programmable, secure ad intent tracking** —  
for a world where impressions don’t mean squat unless they're seen *and* touched.

Admojo NFC Writer:  
Small device. Big proof.


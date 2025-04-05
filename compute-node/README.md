# 🧠 AdMojo Compute Node: Where Eyeballs Meet Blockchain 👁️‍🗨️⛓️

## What Sorcery Is This? 🪄

Welcome, fellow code wizard! You've discovered the magical realm where computer vision and blockchain collide to create undeniable proof that humans actually looked at those ads. Wild, right?

## The Tech Stack of Champions 🏆

- **Go**: Because life's too short for garbage collection pauses
- **OpenCV**: Computer vision that doesn't require a PhD
- **Ethereum**: Because regular databases are _so_ Web2
- **Deep Neural Networks**: Basically magic, but with math

## ✨ Features That Make You Go "Woah" ✨

### 🔍 AI-powered People Spotting
- **Face Detection**: Our SSD + ResNet model can spot a face faster than you can say "privacy concerns"
- **Clever Deduplication**: Using IoU tracking, we ensure we don't count the same bewildered face twice
- **Configurable Paranoia Level**: Adjust confidence thresholds to your trust issues
- **See The Matrix**: Watch the detector work in real-time with cool green rectangles (sunglasses not included)

### 🔗 Blockchain Verification (The Trust Machine)
- **Firmware Fingerprinting**: Each ESP32-CAM device has a unique cryptographic identity
- **Tamper-Evident Metrics**: Try to fake these numbers. Seriously, we dare you
- **Smart Contract Integration**: Your ad metrics are immortalized on the blockchain (until ETH2.0 finally ships)
- **Cryptographic Flex**: ECDSA signatures that would make Satoshi nod in approval

## 🔐 The Verification Dance

1. **Trust, But Verify**: 
   - We pull cryptographic receipts from each camera
   - If the firmware is tampered, the signature won't match
   - Blockchain don't lie, people don't lie, only ads lie

2. **Face-Finding Wizardry**:
   - Every 5 seconds: "Got your face!"
   - DNN goes brrrrr
   - IoU says "Is this the same person as before? 🤔"

3. **The Count of Monte Crypto**:
   - We tally viewers like a vegas dealer counts cards
   - But unlike Vegas, our count is provably fair

4. **Blockchain Immortality**:
   - Your metrics go on-chain faster than you can say "gas fees"
   - Each transaction says: "These many humans saw this ad, for realsies"
   - Cryptographically signed, mathematically sealed, digitally delivered

## 🧩 Smart Contract Spellbook

The `oracle/performance_oracle.go` contains incantations to communicate with the ethereal realm:

```go
// Simplified for mere mortals
perfOracle.UpdateMetricsWithSig(
    transactor,
    deviceID,        // Which camera saw the thing
    timestamp,       // When it saw the thing
    faceCount,       // How many faces it saw
    taps,            // How many people tapped (always 0 for now)
    firmwareHash,    // Cryptographic fingerprint
    signature,       // Magical seal of authenticity
)
```

## 🔮 Future Roadmap (When We Get VC Funding)

- Zero-knowledge proofs (because we actually do care about privacy... kinda)
- Multi-device correlation (creepy, but effective)
- NFT-based ad allocation (because one buzzword isn't enough)
- Decentralized ad marketplace (cutting out the middleman, becoming the middleman)

---

Made with ☕ caffeine, 🤖 AI, and a healthy dose of 🤦‍♂️ debugging
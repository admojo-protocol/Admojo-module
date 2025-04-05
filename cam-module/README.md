# 🔐 AdMojoModule

> Zero-trust camera hardware with crypto-signed firmware for Web3 proof-of-views.

## TL;DR
ESP32-CAM module that streams MJPEG + ECDSA signatures → Go Server → L2 Smart Contract validation. Implements cryptographic proof-of-view for decentralized ad networks.

## Core Stack
- **Hardware**: ESP32-CAM (AI Thinker)
- **Crypto**: secp256k1 ECDSA (same as ETH) + SHA-256
- **Transport**: HTTPS with signature headers
- **Verification**: On-chain via `ecrecover`

## Trust Architecture

```
ESP32-CAM            GoLang Server           Smart Contract (L2)
 │                        │                         │
 ├─ Hash firmware ─┐      │                         │
 │                 │      │                         │
 ├─ Sign w/privkey ┘      │                         │
 │                        │                         │
 ├─ Stream MJPEG + sigs ──┼───> Verify sigs ────────┤
 │                        │                         │
 └────────────────────────┴─────> Store proofs ─────┘
```

## Cryptographic Workflow

1. **Boot**: Calculate SHA-256 of firmware in flash
2. **Sign**: Generate secp256k1 signature (r,s,v) of hash
3. **Stream**: Serve MJPEG with HTTP headers:
   ```
   X-Firmware-Hash: 64-char hex
   X-Signature: 130-char hex (r||s||v)
   ```
4. **Verify**: On-chain using `ecrecover(keccak256(firmwareHash), v, r, s)`

## Quick Deploy

```bash
# Flash, assuming arduino-cli setup
arduino-cli compile --fqbn esp32:esp32:esp32cam AdMojoModule
arduino-cli upload -p /dev/ttyUSB0 --fqbn esp32:esp32:esp32cam AdMojoModule

# Access
curl -v http://192.168.1.43/ # Verify headers
```

## Dev Tips

```javascript
// Verify in JS
const pubKey = web3.eth.accounts.recover(
  web3.utils.sha3(firmwareHash),
  signature
);
```

```solidity
// Verify on-chain
function verifyProofOfView(bytes32 hash, uint8 v, bytes32 r, bytes32 s) public view returns (bool) {
    address signer = ecrecover(hash, v, r, s);
    return trustedDevices[signer];
}
```

## Roadmap

| Now | Next |
|-----|------|
| Flash storage | ATECC608A secure element |
| MJPEG streaming | Snapshot event storage |
| Raw signatures | BLS threshold signatures |
| Basic verification | zk-proofs of on-device AI inference |
| HTTP(S) | L2 oracle integration |

## Why This Rocks

Traditional ad metrics: "Trust me bro"  
AdMojoModule: "Verify my cryptographic proof on-chain"
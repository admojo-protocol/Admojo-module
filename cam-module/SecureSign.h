#ifndef SECURE_SIGN_H
#define SECURE_SIGN_H

#include <Arduino.h>

/**
 * @brief Sign the given data using the private key stored in a secure flash section.
 *        The signature is produced on the secp256k1 curve and is output in the Ethereum-
 *        compatible format (65 bytes: 32-byte r || 32-byte s || 1-byte v).
 *
 * @param data      Pointer to the data to sign.
 * @param data_len  Length of the data in bytes.
 * @param signature Output buffer (must be at least 65 bytes) to store the signature.
 * @param sig_len   On output, will be set to 65.
 * @return true on success, false on failure.
 */
bool signData(const uint8_t* data, size_t data_len, uint8_t* signature, size_t* sig_len);

#endif // SECURE_SIGN_H

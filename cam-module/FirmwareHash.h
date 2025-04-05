#ifndef FIRMWARE_HASH_H
#define FIRMWARE_HASH_H

#include <Arduino.h>

/**
 * @brief Compute the SHA-256 hash of the running firmware (the application partition).
 * @param hash_out Pointer to a 32-byte array where the computed hash will be stored.
 * @return true on success, false on failure.
 */
bool computeFirmwareHash(uint8_t* hash_out);

#endif // FIRMWARE_HASH_H

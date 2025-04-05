#include "FirmwareHash.h"

extern "C" {
  #include "esp_partition.h"
  #include "esp_spi_flash.h"
  #include "mbedtls/sha256.h"
}

bool computeFirmwareHash(uint8_t* hash_out)
{
    // Locate the running application partition (OTA_0 or, if not found, factory)
    const esp_partition_t* partition = esp_partition_find_first(
        ESP_PARTITION_TYPE_APP, ESP_PARTITION_SUBTYPE_APP_OTA_0, NULL
    );
    if (!partition) {
        partition = esp_partition_find_first(
            ESP_PARTITION_TYPE_APP, ESP_PARTITION_SUBTYPE_APP_FACTORY, NULL
        );
        if (!partition) {
            Serial.println("[FirmwareHash] Failed to find application partition.");
            return false;
        }
    }

    mbedtls_sha256_context ctx;
    mbedtls_sha256_init(&ctx);
    if (mbedtls_sha256_starts_ret(&ctx, 0) != 0) {
        mbedtls_sha256_free(&ctx);
        return false;
    }

    const size_t blockSize = 512;
    uint8_t buffer[blockSize];
    size_t remaining = partition->size;
    uint32_t address = partition->address;

    while (remaining > 0) {
        size_t toRead = (remaining > blockSize) ? blockSize : remaining;
        esp_err_t err = spi_flash_read(address, buffer, toRead);
        if (err != ESP_OK) {
            Serial.printf("[FirmwareHash] spi_flash_read failed: 0x%x\n", err);
            mbedtls_sha256_free(&ctx);
            return false;
        }
        if (mbedtls_sha256_update_ret(&ctx, buffer, toRead) != 0) {
            mbedtls_sha256_free(&ctx);
            return false;
        }
        remaining -= toRead;
        address   += toRead;
    }

    if (mbedtls_sha256_finish_ret(&ctx, hash_out) != 0) {
        mbedtls_sha256_free(&ctx);
        return false;
    }
    mbedtls_sha256_free(&ctx);
    return true;
}

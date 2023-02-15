// NumKey
//
// countrykey.h
//
// @category   Libraries
// @author     Nicola Asuni <nicola.asuni@vonage.com>
// @copyright  2019-2022 Vonage
// @license    see LICENSE file
// @link       https://github.com/Vonage/numkey

/**
 * @file countrykey.h
 * @brief CountryKey functions.
 *
 *  The functions provided here allows to generate and process a 16 bit Unsigned Integer Keys for ISO 3166 alpha-2 country code.
 */

#ifndef NUMKEY_COUNTRYKEY_H
#define NUMKEY_COUNTRYKEY_H

#include <inttypes.h>
#include <stddef.h>

/**
 * Encode country as numerical key.
 *
 * @param country ISO 3166 alpha-2 country code.
 *
 * @return Encoded country code.
 */
static inline uint16_t countrykey(const char *country)
{
    return ((country[0] << 8) | country[1]);
}

/**
 * Decode countrykey into 2-byte string.
 *
 * @param ck       CountryKey.
 * @param country  Pre-allocated string buffer to be returned (it must be at least two bytes).
 */
static inline void decode_countrykey(uint16_t ck, char *country)
{
    country[0] = (char)((ck & 0xFF00) >> 8);
    country[1] = (char)(ck & 0x00FF);
    country[2] = 0;
}

#endif  // NUMKEY_COUNTRYKEY_H

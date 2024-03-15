// NumKey
//
// prefixkey.h
//
// @category   Libraries
// @author     Nicola Asuni
// @license    see LICENSE file
// @link       https://github.com/Vonage/numkey

/**
 * @file prefixkey.h
 * @brief PrefixKey functions.
 *
 * The functions provided here allows to generate and process a 64 bit Unsigned Integer Keys for numbers up to 18 digits.
 */

#ifndef NUMKEY_PREFIXKEY_H
#define NUMKEY_PREFIXKEY_H

#include <inttypes.h>
#include <stddef.h>

#define PKNUMMAXLEN 15 //!< Maximum number of digits to store for the prefixkey.

/**
 * Encode a number string into uint64.
 * The encoded number is always 15 digits long as it is either right-padded with zeros or truncated.
 * The prefixkey is safe to cast as int64 as it is always smaller than max int64.
 *
 * @param number String containing the E.164 number or prefix (max 15 digits or it will be truncated).
 * @param size   Length of the number (number of digits).
 *
 * @return Encoded number
 */
static inline uint64_t prefixkey(const char *number, size_t size)
{
    uint64_t num = 0;
    uint8_t b;
    size_t i = 0;
    if (size > PKNUMMAXLEN)
    {
        size = PKNUMMAXLEN; // truncate number
    }
    for (i = 0; i < size; i++)
    {
        b = (uint8_t)number[i] - '0';
        num = (num * 10) + b;
    }
    for (i = size; i < PKNUMMAXLEN; i++)
    {
        num = (num * 10); // zero right-padding
    }
    return num;
}

#endif  // NUMKEY_PREFIXKEY_H

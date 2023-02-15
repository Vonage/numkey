// NumKey
//
// numkey.h
//
// @category   Libraries
// @author     Nicola Asuni <nicola.asuni@vonage.com>
// @copyright  2019-2022 Vonage
// @license    see LICENSE file
// @link       https://github.com/Vonage/numkey

/**
 * @file numkey.h
 * @brief NumKey main functions.
 *
 * The functions provided here allows to generate and process a 64 bit Unsigned Integer Keys for Short Codes and E.164 LVN numbers.
 * The NumKey is sortable country and number and it is fully reversible.
 * It can be used to sort, search and match number-based data easily and very quickly.
 */

#ifndef NUMKEY_NUMKEY_H
#define NUMKEY_NUMKEY_H

#include <inttypes.h>
#include <stddef.h>
#include "hex.h"

#define NKBMASK_COUNTRY_FL 0xF800000000000000  //!< Bit mask for the ISO 3166 alpha-2 country code first letter  [ 11111000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 ].
#define NKBMASK_COUNTRY_SL 0x07C0000000000000  //!< Bit mask for the ISO 3166 alpha-2 country code second letter [ 00000111 11000000 00000000 00000000 00000000 00000000 00000000 00000000 ].
#define NKBMASK_NUMBER     0x003FFFFFFFFFFFF0  //!< Bit mask for the short code or E.164 number (max 15 digits)  [ 00000000 00111111 11111111 11111111 11111111 11111111 11111111 11110000 ].
#define NKBMASK_LENGTH     0x000000000000000F  //!< Bit mask for the number length                               [ 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00001111 ].

#define NKBSHIFT_COUNTRY_FL 59 //!< COUNTRY first letter LSB position from the NumKey LSB.
#define NKBSHIFT_COUNTRY_SL 54 //!< COUNTRY second letter LSB position from the NumKey LSB.
#define NKBSHIFT_NUMBER      4 //!< NUMBER LSB position from the NumKey LSB.

#define NKSLENGTH_COUNTRY    3 //!< Number of characters in the country code + NULL terminator.
#define NKSLENGTH_NUMBER    16 //!< Number of characters in the number code + NULL terminator.

#define NKCSHIFT_CHAR       64 //!< Value shift to encode characters to numbers (A=1, ..., Z=26).

#define NKNUMMAXLEN         15 //!< Maximum number length for E.164 and key reversibility.

/**
 * NumKey struct.
 * Contains the NumKey components (COUNTRY, NUMBER).
 */
typedef struct numkey_t
{
    char country[NKSLENGTH_COUNTRY]; //!< ISO 3166 alpha-2 country code.
    char number[NKSLENGTH_NUMBER];   //!< Short code or E.164 number (max 15 digits).
} numkey_t;

/**
 * Encodes the input character to a numeric value.
 * NOTE: This is safe to be used only with A to Z characters.
 *
 * @param c Character to encode.
 *
 * @return Encoded character.
 */
static inline uint64_t encode_char(int c)
{
    return (uint64_t)(c - NKCSHIFT_CHAR);
}

/**
 * Encode country code for numkey.
 *
 * @param country ISO 3166 alpha-2 country code.
 *
 * @return Encoded country code.
 */
static inline uint64_t encode_country(const char *country)
{
    return ((encode_char(country[0]) << NKBSHIFT_COUNTRY_FL) | (encode_char(country[1]) << NKBSHIFT_COUNTRY_SL));
}

/**
 * Decode country code into 2-byte string.
 *
 * @param nk       NumKey.
 * @param country  Pre-allocated string buffer to be returned (it must be at least two bytes).
 */
static inline void decode_country(uint64_t nk, char *country)
{
    country[0] = (char)((nk & NKBMASK_COUNTRY_FL) >> NKBSHIFT_COUNTRY_FL) + NKCSHIFT_CHAR;
    country[1] = (char)((nk & NKBMASK_COUNTRY_SL) >> NKBSHIFT_COUNTRY_SL) + NKCSHIFT_CHAR;
    country[2] = 0;
}

/**
 * Encode number string.
 *
 * @param number String containing the Short code or E.164 LVN number.
 * @param size   Length of the number (number of digits).
 *
 * @return Encoded number
 */
static inline uint64_t encode_number(const char *number, size_t size)
{
    uint64_t num = 0;
    uint8_t b;
    size_t i, j = 0;
    uint64_t len = (uint64_t)(size);
    if (size > NKNUMMAXLEN)
    {
        j = (size - NKNUMMAXLEN); // last 15 digits
        len = 0;                  // flag non-revesible encoding
    }
    for (i = j; i < size; i++)
    {
        b = (uint8_t)number[i] - '0';
        num = (num * 10) + b;
    }
    return ((num << NKBSHIFT_NUMBER) | (len & NKBMASK_LENGTH));
}

/**
 * Decode number into string.
 *
 * @param nk       NumKey.
 * @param number   Number string buffer to be returned.
 *
 * @return      Number length (number of digits).
 */
static inline size_t decode_number(uint64_t nk, char *number)
{
    size_t size = (size_t)(nk & NKBMASK_LENGTH);
    uint64_t num = ((nk & NKBMASK_NUMBER) >> NKBSHIFT_NUMBER);
    int rem = 0;
    int i;
    for (i = (size - 1); i >= 0; i--)
    {
        rem = (num % 10);
        number[i] = (rem + '0');
        num = (num / 10);
    }
    number[size] = 0;
    return size;
}

/**
 * Encode numkey.
 *
 * @param country ISO 3166 alpha-2 country code.
 * @param number  String containing the Short code or LVN number.
 * @param numsize Length of the number (number of digits).
 *
 * @return NumKey 64 bit code.
 */
static inline uint64_t numkey(const char *country, const char *number, size_t numsize)
{
    return (encode_country(country) | encode_number(number, numsize));
}

/**
 * Decode a NumKey code to get the individual components.
 *
 * @param nk    NumKey code.
 * @param data  Decoded numkey structure.
 */
static inline void decode_numkey(uint64_t nk, numkey_t *data)
{
    decode_country(nk, data->country);
    decode_number(nk, data->number);
}

static inline int8_t compare_uint64_t(uint64_t a, uint64_t b)
{
    return (a < b) ? -1 : (a > b);
}

/**
 * Compares two NumKeys by country only.
 *
 * @param nka    The first NumKey to be compared.
 * @param nkb    The second NumKey to be compared.
 *
 * @return -1 if the first country is smaller than the second, 0 if they are equal and 1 if the first is greater than the second.
 */
static inline int8_t compare_numkey_country(uint64_t nka, uint64_t nkb)
{
    return compare_uint64_t((nka >> NKBSHIFT_COUNTRY_SL), (nkb >> NKBSHIFT_COUNTRY_SL));
}

/**
 * Returns NumKey hexadecimal string (16 characters).
 *
 * @param nk    NumKey code.
 * @param str   String buffer to be returned (it must be sized 17 bytes at least).
 *
 * @return      Upon successful return, these function returns the number of characters processed
 *              (excluding the null byte used to end output to strings).
 *              If the buffer size is not sufficient, then the return value is the number of characters required for
 *              buffer string, including the terminating null byte.
 */
static inline size_t numkey_hex(uint64_t nk, char *str)
{
    return hex_uint64_t(nk, str);
}

/**
 * Parses a NumKey hexadecimal string and returns the code.
 *
 * @param ns NumKey hexadecimal string (it must contain 16 hexadecimal characters).
 *
 * @return A NumKey code.
 */
static inline uint64_t parse_numkey_hex(const char *ns)
{
    return parse_hex_uint64_t(ns);
}

#endif  // NUMKEY_NUMKEY_H

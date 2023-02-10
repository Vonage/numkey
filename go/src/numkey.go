// Package numkey provides 64 bit Encoding for Short Codes and E.164 LVN.
//
// @category   Libraries
// @author     Nicola Asuni <nicola.asuni@vonage.com>
// @copyright  2019-2020 Vonage
// @license    see LICENSE file
// @link       https://github.com/vonage/numkey
package numkey

import (
	"strconv"
	"strings"
)

// NkbmaskCountryFl is the bit mask for the ISO 3166 alpha-2 country code first letter  [ 11111000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 ]
const NkbmaskCountryFl = 0xF800000000000000

// NkbmaskCountrySl is it mask for the ISO 3166 alpha-2 country code second letter [ 00000111 11000000 00000000 00000000 00000000 00000000 00000000 00000000 ]
const NkbmaskCountrySl = 0x07C0000000000000

// NkbmaskNumber is the bit mask for the short code or E.164 number (max 15 digits) [ 00000000 00111111 11111111 11111111 11111111 11111111 11111111 11110000 ]
const NkbmaskNumber = 0x003FFFFFFFFFFFF0

// NkbmaskLength is the bit mask for the number length [ 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00001111 ]
const NkbmaskLength = 0x000000000000000F

// NkbshiftCountryFl is the country first letter LSB position from the NumKey LSB
const NkbshiftCountryFl = 59

// NkbshiftCountrySl is the country second letter LSB position from the NumKey LSB
const NkbshiftCountrySl = 54

// NkbshiftNumber is the number LSB position from the NumKey LSB
const NkbshiftNumber = 4

// NkcshiftChar is value shift to encode characters to numbers (A=1, ..., Z=26)
const NkcshiftChar = 64

// NumMaxLen is the maximum number length for E.164 and key reversibility
const NumMaxLen = 15

// encodeChar encodes the input character to a numeric value.
// NOTE: This is safe to be used only with A to Z characters.
func encodeChar(c byte) uint64 {
	return uint64(int(c) - NkcshiftChar)
}

// encodeCountry encodes the country code.
// The country parameter must be an ISO 3166 alpha-2 country code.
// Returns 0 in case of error
func encodeCountry(country string) uint64 {
	return ((encodeChar(country[0]) << NkbshiftCountryFl) | (encodeChar(country[1]) << NkbshiftCountrySl))
}

// decodeCountry decodes the country code into 2-byte string.
func decodeCountry(nk uint64) string {
	return string([]byte{
		byte(((nk & NkbmaskCountryFl) >> NkbshiftCountryFl) + NkcshiftChar),
		byte(((nk & NkbmaskCountrySl) >> NkbshiftCountrySl) + NkcshiftChar),
	})
}

// encodeNumber encodes a number string (shortcode or E.164).
// Numbers are left-truncated to fit 15 digits and the length is set to 16.
func encodeNumber(number string) uint64 {
	size := len(number)
	if size > NumMaxLen {
		number = number[(size - NumMaxLen):size] // last 15 digits
		size = 0                                 // flag non-revesible encoding
	}
	num, err := strconv.ParseUint(number, 10, 64)
	if err != nil {
		return 0
	}
	return ((num << NkbshiftNumber) | ((uint64)(size) & NkbmaskLength))
}

// func decodeNumber decodes a number into a string
func decodeNumber(nk uint64) string {
	size := int(nk & NkbmaskLength)
	if size == 0 {
		return "" // non-reversible number encoding
	}
	s := strconv.FormatUint(((nk & NkbmaskNumber) >> NkbshiftNumber), 10)
	slen := len(s)
	if slen < size {
		return strings.Repeat("0", (size-slen)) + s
	}
	return s
}

// NumKey returns an encoded COUNTRY + NUMBER
func NumKey(country, number string) uint64 {
	size := len(number)
	if len(country) != 2 || size < 1 {
		return 0
	}
	return (encodeCountry(country) | encodeNumber(number))
}

// DecodeNumKey parses a numkey string and returns the components as TNumKey structure.
// Return country and number strings.
func DecodeNumKey(nk uint64) (string, string) {
	if nk == 0 {
		return "", ""
	}
	return decodeCountry(nk), decodeNumber(nk)
}

func compareUint64(a, b uint64) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// CompareNumKeyCountry compares two NumKeys by country only.
func CompareNumKeyCountry(nka, nkb uint64) int {
	return compareUint64((nka >> NkbshiftCountrySl), (nkb >> NkbshiftCountrySl))
}

// Hex provides a 16 digits hexadecimal string representation of a 64bit unsigned number.
func Hex(v uint64) string {
	s := strconv.FormatUint(v, 16)
	slen := len(s)
	if slen < 16 {
		return strings.Repeat("0", (16-slen)) + s
	}
	return s
}

// ParseHex parses a 16 digit HEX string and returns the 64 bit unsigned number.
// return 0 in case of error
func ParseHex(s string) uint64 {
	num, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return 0
	}
	return num
}

// Package numkey is a Go wrapper for the numkey C software library.
// 64 bit Encoding for Short Codes and E.164 LVN.
//
// @category   Libraries
// @author     Nicola Asuni
// @license    see LICENSE file
// @link       https://github.com/Vonage/numkey
package numkey

/*
#cgo CFLAGS: -O3 -pedantic -std=c99 -Wextra -Wno-strict-prototypes -Wcast-align -Wundef -Wformat -Wformat-security -Wshadow
#include <stdlib.h>
#include <inttypes.h>
#include "../../c/src/numkey/hex.h"
#include "../../c/src/numkey/numkey.h"
#include "../../c/src/numkey/prefixkey.h"
#include "../../c/src/numkey/countrykey.h"
*/
import "C"
import "unsafe"

// TNumKey contains the number components.
type TNumKey struct {
	Country string `json:"country"`
	Number  string `json:"number"`
}

// castCNumKey convert C numkey_t to GO TNumKey.
func castCNumKey(nk C.numkey_t) TNumKey {
	return TNumKey{
		Country: C.GoString((*C.char)(unsafe.Pointer(&nk.country[0]))),
		Number:  C.GoString((*C.char)(unsafe.Pointer(&nk.number[0]))),
	}
}

// StringToNTBytesN convert a string to byte array allocating "size" bytes.
func StringToNTBytesN(s string, size int) []byte {
	b := make([]byte, size)
	copy(b, s)

	return b
}

// NumKey returns an encoded COUNTRY + NUMBER.
// If the country or number are invalid this function returns 0.
func NumKey(country, number string) uint64 {
	countrysize := len(country)
	numsize := len(number)

	if countrysize != 2 || numsize < 1 {
		return 0
	}

	bcountry := StringToNTBytesN(country, countrysize+1)
	bnumber := StringToNTBytesN(number, numsize+1)
	pcountry := unsafe.Pointer(&bcountry[0]) // #nosec
	pnumber := unsafe.Pointer(&bnumber[0])   // #nosec

	return uint64(C.numkey((*C.char)(pcountry), (*C.char)(pnumber), C.size_t(numsize)))
}

// DecodeNumKey parses a numkey string and returns the components as TNumKey structure.
func DecodeNumKey(nk uint64) TNumKey {
	if nk == 0 {
		return TNumKey{}
	}

	var data C.numkey_t

	C.decode_numkey(C.uint64_t(nk), &data)

	return castCNumKey(data)
}

// CompareNumKeyCountry compares two NumKeys by country only.
func CompareNumKeyCountry(nka, nkb uint64) int {
	return int(C.compare_numkey_country(C.uint64_t(nka), C.uint64_t(nkb)))
}

// Hex provides a 16 digits hexadecimal string representation of a 64bit unsigned number.
func Hex(v uint64) string {
	cstr := C.malloc(17)

	defer C.free(cstr)

	C.numkey_hex(C.uint64_t(v), (*C.char)(cstr))

	return C.GoStringN((*C.char)(cstr), C.int(16))
}

// ParseHex parses a 16 digit HEX string and returns the 64 bit unsigned number.
func ParseHex(s string) uint64 {
	b := StringToNTBytesN(s, len(s)+1)
	p := unsafe.Pointer(&b[0]) // #nosec

	return uint64(C.parse_numkey_hex((*C.char)(p)))
}

// PrefixKey encodes a number string into uint64.
// The encoded number is always 15 digits long as it is either right-padded with zeros or truncated.
// The prefixkey is safe to cast as int64 as it is always smaller than max int64.
func PrefixKey(number string) uint64 {
	numsize := len(number)
	bnumber := StringToNTBytesN(number, numsize+1)
	pnumber := unsafe.Pointer(&bnumber[0]) // #nosec

	return uint64(C.prefixkey((*C.char)(pnumber), C.size_t(numsize)))
}

// CountryKey encodes ISO 3166 alpha-2 country code into uint16.
func CountryKey(country string) uint16 {
	countrysize := len(country)
	if countrysize != 2 {
		return 0
	}

	bcountry := StringToNTBytesN(country, countrysize+1)
	pcountry := unsafe.Pointer(&bcountry[0]) // #nosec

	return uint16(C.countrykey((*C.char)(pcountry)))
}

// DecodeCountryKey decodes countrykey into ISO 3166 alpha-2 country code.
func DecodeCountryKey(ck uint16) string {
	cstr := C.malloc(3)

	defer C.free(cstr)

	C.decode_countrykey(C.uint16_t(ck), (*C.char)(cstr))

	return C.GoStringN((*C.char)(cstr), C.int(2))
}

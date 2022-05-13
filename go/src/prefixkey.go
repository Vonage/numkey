package numkey

import (
	"strconv"
)

// PrefixMaxLen is the maximum number of digits to store for the prefixkey.
const PrefixMaxLen = 15

// PrefixKey encodes a number string into uint64.
// The encoded number is always 15 digits long as it is either right-padded with zeros or truncated.
// The prefixkey is safe to cast as int64 as it is always smaller than max int64.
func PrefixKey(number string) uint64 {
	var num uint64

	size := len(number)

	if size > PrefixMaxLen {
		size = PrefixMaxLen
		number = number[0:PrefixMaxLen] // truncate number
	}

	num, err := strconv.ParseUint(number, 10, 64)
	if err != nil {
		return 0
	}

	for i := size; i < PrefixMaxLen; i++ {
		num = (num * 10) // zero right-padding
	}

	return num
}

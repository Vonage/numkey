package main

// Usage example for NumKey GO version
// https://github.com/Vonage/numkey

import (
	"fmt"

	nk "github.com/Vonage/numkey/go/src"
)

func main() {

	// BASIC NUMKEY FUNCTIONS
	// --------------------------

	fmt.Println(nk.DecodeNumKey(14027409114588157055))
	// {XJ 762942138198343}

	fmt.Println(nk.NumKey("XJ", "762942138198343"))
	// 14027409114588157055

	fmt.Println(nk.CompareNumKeyCountry(0xd6a23089b8e15cdf, 0xd6a2300000000000))
	// 0

	fmt.Println(nk.Hex(14027409114588157055))
	// c2ab5e44f21a947f

	fmt.Println(nk.ParseHex("c2ab5e44f21a947f"))
	// 14027409114588157055
}

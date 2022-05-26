package numkey

// CountryKey encodes ISO 3166 alpha-2 country code into uint16.
func CountryKey(country string) uint16 {
	return (uint16(country[0]) << 8) | uint16(country[1])
}

// DecodeCountryKey decodes countrykey into ISO 3166 alpha-2 country code.
func DecodeCountryKey(ck uint16) string {
	return string([]byte{
		byte((ck & 0xFF00) >> 8),
		byte(ck & 0x00FF),
	})
}

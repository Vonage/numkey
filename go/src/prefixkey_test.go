package numkey

// prefixkey_test.go
// @category   Libraries
// @author     Nicola Asuni
// @license    see LICENSE file
// @link       https://github.com/Vonage/numkey

import "testing"

// TPrefixKeyData contains test data.
type TPrefixKeyData struct {
	number string
	pk     uint64
}

func prefixkeyTestData() []TPrefixKeyData {
	return []TPrefixKeyData{
		{"", 0},
		{"0", 0},
		{"00", 0},
		{"00000000000000", 0},
		{"000000000000000", 0},
		{"0000000000000000", 0},
		{"000000000000001", 1},
		{"0000000000000019", 1},
		{"1", 100000000000000},
		{"10", 100000000000000},
		{"10000000000000", 100000000000000},
		{"100000000000000", 100000000000000},
		{"1000000000000000", 100000000000000},
		{"999999999999999", 999999999999999},
		{"abcdef", 0},
	}
}

func TestPrefixKey(t *testing.T) {
	t.Parallel()

	for _, v := range prefixkeyTestData() {
		t.Run("", func(t *testing.T) {
			t.Parallel()

			pk := PrefixKey(v.number)
			if pk != v.pk {
				t.Errorf("The code value is different, expected %#v got %#v", v.pk, pk)
			}
		})
	}
}

func BenchmarkPrefixKey(b *testing.B) {
	b.ResetTimer()

	for range b.N {
		PrefixKey("123456789012345")
	}
}

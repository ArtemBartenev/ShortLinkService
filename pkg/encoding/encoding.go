package encoding

import (
	"math"
	"math/rand"
)

const (
	ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	BASE     = len(ALPHABET)
)

func GenerateStringByBase62() string {
	var symbols []byte
	number := uint(rand.Uint32())

	for number > 0 {
		remainder := math.Mod(float64(number), float64(BASE))
		symbol := ALPHABET[int(remainder)]
		symbols = append(symbols, symbol)

		number /= uint(BASE)
	}

	return string(symbols)
}

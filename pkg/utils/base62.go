package utils
import (
	"bytes"
"math"
)

const table = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func EncodeBase62(number int64) string {
	if number == 0 {
		return string(table[0])
	}

	chars := make([]byte, 0)

	length := int64(len(table))

	for number > 0 {
		result    := number / length
		remainder := number % length
		chars   = append(chars, table[remainder])
		number  = result
	}

	for i, j := 0, len(chars) - 1; i < j; i, j = i + 1, j - 1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}

func DecodeBase62(token string) int64 {
	number := int64(0)
	idx    := 0.0
	chars  := []byte(table)

	charsLength := float64(len(chars))
	tokenLength := float64(len(token))

	for _, c := range []byte(token) {
		power := tokenLength - (idx + 1)
		index := int64(bytes.IndexByte(chars, c))
		number += index * int64(math.Pow(charsLength, power))
		idx++
	}

	return number
}

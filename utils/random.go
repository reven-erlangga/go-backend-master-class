package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generate a random integer between min and max number
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min + 1)
}

// RandomString generate a random string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < k; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(8)
}

func RandomMoney() int64 {
	return RandomInt(0, 800)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "IDR"}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}
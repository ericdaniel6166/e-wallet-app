package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt returns a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)

}

// RandomString returns a random string with length n
func RandomString(n int) string {
	var strBuilder strings.Builder
	lenAlphabet := len(alphabet)
	for i := 0; i < n; i++ {
		char := alphabet[rand.Intn(lenAlphabet)]
		strBuilder.WriteByte(char)
	}
	return strBuilder.String()
}

// RandomOwner returns a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	lenCurrencies := len(currencies)
	return currencies[rand.Intn(lenCurrencies)]
}

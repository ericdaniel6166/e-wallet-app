package util

import (
	"github.com/shopspring/decimal"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt returns a random integer in the half-open interval [min,max)
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

// RandomFloat returns a random float in the half-open interval [min,max)
func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
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

func RandomUsername() string {
	return RandomString(12)
}

// RandomMoney generates a random amount of money
func RandomMoney() decimal.Decimal {
	return decimal.NewFromFloat(RandomFloat(1000, 10000))
}

package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyz")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return int64(RandomInt(0, 1000))
}

// RandomCurrency generates a random currency
func RandomCurrency() string {
	currencies := []string{USD, EUR, NGN, GBP}
	return currencies[rand.Intn(len(currencies))]
}


func RandomEmail() string {
	return fmt.Sprintf("%s@payzonex.com", Ran)
}
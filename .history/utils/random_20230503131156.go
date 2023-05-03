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
	return fmt.Sprintf("%s@payzonex.com", RandomString(6))
}

func SplitStrings(s string) []string {
	var r []string
	for _, v := range s {
		r = append(r, string(v))
	}
	return r
}

func RandomPhoneNumber() string {
	// Nigerian Phone Number 
	// should start with 0 and have 11 digits
	// e.g 08012345678
	int := RandomInt(10000000, 9999999)
	return fmt.Sprintf("080%s", RandomInt(10000000, 99999999))
}
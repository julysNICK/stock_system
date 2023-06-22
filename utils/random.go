package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := (alphabet[rand.Intn(k)])
		sb.WriteByte(c)
	}
	return sb.String()

}

func RandomInt64(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

func RandomName() string {
	return RandomString(6)
}

func RandomEmail() string {
	return RandomString(6) + "@gmail.com"
}

func RandomAddress() string {
	return RandomString(6)
}

func RandomPhone() string {
	return RandomString(6)
}

func RandomPassword() string {
	return RandomString(10)
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandomDate() time.Time {
	return time.Now().AddDate(0, 0, RandomInt(1, 100))
}

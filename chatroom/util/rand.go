package util

import "math/rand"

var charRunes = []rune("0123456789abcdef")

// RandStringID returns a random string of size n which is composed of digits from 0-9.
func RandStringID(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = charRunes[rand.Intn(len(charRunes))]
	}

	return string(b)
}

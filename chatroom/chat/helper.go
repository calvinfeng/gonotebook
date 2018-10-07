package chat

import (
	"fmt"
	"math/rand"
	"time"
)

var charRunes = []rune("0123456789abcdef")

// randStringID returns a random string of size n which is composed of digits from 0-9.
func randStringID(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = charRunes[rand.Intn(len(charRunes))]
	}

	return string(b)
}

func loginfo(msg string) {
	fmt.Printf("[INFO][%s] %s\n", time.Now(), msg)
}

func logerr(trace string, err error) {
	fmt.Printf("[EROR][%s] %s - %s\n", time.Now(), trace, err.Error())
}

package katoptron

import (
	"crypto/aes"
	"math/rand"
	"testing"
	"time"
)

func key() string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits + specials
	length := 32
	buffer := make([]byte, length)
	for i := 0; i < length; i++ {
		buffer[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buffer), func(i, j int) {
		buffer[i], buffer[j] = buffer[j], buffer[i]
	})
	return string(buffer)
}

func TestPkgType(t *testing.T) {
	k := key()
	cipher, err := aes.NewCipher([]byte(k))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(Any(cipher))
	Display("hihi", cipher)
}

func TestValueType(t *testing.T) {
	now := time.Now()

	t.Log(Any(now))
}

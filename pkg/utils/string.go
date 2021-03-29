package utils

import (
	"crypto/rand"
	"math/big"
)

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(defaultLetters))))
		b[i] = defaultLetters[n.Int64()]
	}
	return string(b)
}

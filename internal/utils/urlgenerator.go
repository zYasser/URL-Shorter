package utils

import (
	"math/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateShortURL() string {
	url := make([]byte, 7)
	for i := 0; i < 7; i++ {
		url[i] = letters[rand.Intn(len(letters))]
	}

	return string(url)
}

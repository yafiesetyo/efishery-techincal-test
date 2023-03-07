package str

import (
	"math/rand"
	"time"
)

func GeneratePassword() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().Unix())

	length := 4

	randStr := make([]rune, length)
	for i := 0; i < length; i++ {
		randStr[i] = letters[rand.Intn(len(letters))]
	}

	return string(randStr)
}

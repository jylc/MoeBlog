package utils

import (
	"math/rand"
	"time"
)


func GenRandomString(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	s := make([]rune, n)

	for i := range s {
		s[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	return string(s)
}

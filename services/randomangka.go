package services

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(length int) string {
	// Mengatur seed acak berdasarkan waktu saat ini
	rand.Seed(time.Now().UnixNano())

	// Daftar karakter yang mungkin
	chars := "0123456789"

	// Membangun string acak dengan panjang yang ditentukan
	randomNumber := make([]byte, length)
	for i := 0; i < length; i++ {
		randomNumber[i] = chars[rand.Intn(len(chars))]
	}

	return string(randomNumber)
}

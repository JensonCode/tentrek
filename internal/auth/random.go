package auth

import (
	"math/rand"
	"time"
)

const (
	charBytes    = "abcdefghijklmnopqrstuvwxyz"
	capCharBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numBytes     = "0123456789"
	specialBytes = "!@#$%^&*()_+{}[]|\\;:'<>,.?/~`"
)

func GenerateRandomPassword() string {
	rand.NewSource(time.Now().UnixNano())

	password := make([]byte, 32)

	for i := range password {
		switch rand.Intn(4) {
		case 0:
			password[i] = charBytes[rand.Intn(len(charBytes))]
		case 1:
			password[i] = capCharBytes[rand.Intn(len(capCharBytes))]
		case 2:
			password[i] = numBytes[rand.Intn(len(numBytes))]
		case 3:
			password[i] = specialBytes[rand.Intn(len(specialBytes))]
		}
	}

	return string(password)
}

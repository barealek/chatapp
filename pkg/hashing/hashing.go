package hashing

import (
	"crypto/sha256"
)

func HashString(in string) []byte {
	hashedPassword := sha256.New()
	hashedPassword.Write([]byte(in))

	return hashedPassword.Sum(nil)
}

package user_db

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

const (
	timeCost   = 4
	memoryCost = 64 * 1024
	threads    = 4
	hashLength = 32
	saltLength = 32
)

// return (hash, salt).
func hashNewPassword(password string) ([]byte, []byte) {
	salt := make([]byte, saltLength)

	controlNum, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}

	if controlNum != saltLength {
		panic("salt length mismatch")
	}

	hashedPassword := hashPassword(password, salt)

	return hashedPassword, salt
}

func hashPassword(password string, salt []byte) []byte {
	hashedPassword := argon2.IDKey(
		[]byte(password),
		salt,
		timeCost,
		memoryCost,
		threads,
		hashLength,
	)

	return hashedPassword
}

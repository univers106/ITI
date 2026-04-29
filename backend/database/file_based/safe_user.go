package file_based

import (
	"crypto/rand"
	"slices"

	"github.com/univers106/ITI/database"
	"golang.org/x/crypto/argon2"
)

const (
	argon2TimeCost     = 1
	argon2MemoryCost   = 64 * 1024
	argon2KeyLength    = 32
	argon2SaltByteSize = 16
	argon2Parallelism  = 4
)

type safeUser struct {
	database.User

	PasswordHash []byte `json:"passwordHash"`
	PasswordSalt []byte `json:"passwordSalt"`
}

func (s *safeUser) SetPassword(password string) {
	s.PasswordSalt = make([]byte, argon2SaltByteSize)
	rand.Read(s.PasswordSalt)
	s.PasswordHash = argon2.IDKey(
		[]byte(password),
		s.PasswordSalt,
		argon2TimeCost,
		argon2MemoryCost,
		argon2Parallelism,
		argon2KeyLength,
	)
}

func (s *safeUser) CheckPassword(password string) bool {
	hash := argon2.IDKey(
		[]byte(password),
		s.PasswordSalt,
		argon2TimeCost,
		argon2MemoryCost,
		argon2Parallelism,
		argon2KeyLength,
	)

	return slices.Equal(s.PasswordHash, hash)
}

func (s *safeUser) GetUser() *database.User {
	return &s.User
}

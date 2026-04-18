package filebased

import (
	"crypto/rand"
	"slices"

	"github.com/univers106/ITI/database"
	"golang.org/x/crypto/argon2"
)

type safeUser struct {
	database.User
	PasswordHash []byte "json:\"passwordHash\""
	PasswordSalt []byte "json:\"passwordSalt\""
}

func (s *safeUser) SetPassword(password string) {
	s.PasswordSalt = make([]byte, 16)
	rand.Read(s.PasswordSalt)
	s.PasswordHash = argon2.IDKey([]byte(password), s.PasswordSalt, 1, 64*1024, 4, 32)

}

func (s *safeUser) CheckPassword(password string) bool {
	hash := argon2.IDKey([]byte(password), s.PasswordSalt, 1, 64*1024, 4, 32)
	return slices.Equal(s.PasswordHash, hash)
}

func (s *safeUser) GetUser() *database.User {
	return &s.User
}

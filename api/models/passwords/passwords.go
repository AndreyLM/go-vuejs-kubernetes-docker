package passwords

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	salt = 10
)

// Encrypt - encryptes password
func Encrypt(pass string) (string, error) {
	val, err := bcrypt.GenerateFromPassword([]byte(pass), salt)
	return string(val), err
}

// Compare - check in password is valid
func Compare(hash, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		return false
	}

	return true
}

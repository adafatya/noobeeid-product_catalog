package utils

import "golang.org/x/crypto/bcrypt"

// encrypt password using bcrypt
// param password string
// return hashed password string, error
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// compare password and hash password
// param password string, hashed password string
// return bool
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

package util

import (
	"math/rand"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := RandomString(10)

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("Unexpected error hashing password: %v", err)
	}

	if hashedPassword == "" {
		t.Error("Hashed password should not be empty")
	}

	err = ComparePasswords(password, hashedPassword)
	if err != nil {
		t.Errorf("Password comparison failed: %v", err)
	}
}

// RandomString generates a random string of the given length.
func RandomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

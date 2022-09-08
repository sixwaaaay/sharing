package secu

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
	"time"
)

// characters
const (
	characters string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#%^&*()-_ []{}<>~`+=,.;:/?|"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenSalt function to generate a random n-digit password
func GenSalt(n int) string {
	builder := strings.Builder{}
	for i := 0; i < n; i++ {
		builder.WriteByte(characters[rand.Intn(len(characters))])
	}
	return builder.String()
}

// Hash function to hash a password (bcrypt algo)
func Hash(password string) string {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(fromPassword)
}

func Compare(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func GenHashedPassAndSalt(password string) (string, string) {
	salt := GenSalt(6)
	hash := Hash(password)
	return hash, salt
}

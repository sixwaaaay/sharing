package secu

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestGenSalt(t *testing.T) {
	salts := map[string]struct{}{}
	for i := 0; i < 100; i++ {
		salt := GenSalt(6)
		if len(salt) != 6 {
			t.Errorf("salt length error: %d", len(salt))
		}
		if _, ok := salts[salt]; ok {
			t.Errorf("duplicate salt: %s", salt)
		}
		salts[salt] = struct{}{}
	}
}

func TestHash(t *testing.T) {
	salt := GenSalt(6)
	password := "123456"
	hash := Hash(password, salt)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt))
	if err != nil {
		t.Errorf("hash compare error: %s", err)
	}
}
func TestCompare(t *testing.T) {
	salt := GenSalt(6)
	password := "123456"
	hash := Hash(password, salt)
	if Compare(password, salt, hash) != true {
		t.Errorf("hash compare error: %s", "password not match")
	}
}

func TestGenHashedPassAndSalt(t *testing.T) {
	password := "123456"
	hash, salt := GenHashedPassAndSalt(password)
	compare := Compare(salt, password, hash)
	require.Equal(t, true, compare)
}

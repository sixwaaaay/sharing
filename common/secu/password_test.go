package secu

import (
	"fmt"
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
	password := "123456"
	hash := Hash(password)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		t.Errorf("hash compare error: %s", err)
	}
}
func TestCompare(t *testing.T) {
	password := "123456"
	hash := Hash(password)
	if Compare(password, hash) != true {
		t.Errorf("hash compare error: %s", "password not match")
	}
}

func TestGenHashedPassAndSalt(t *testing.T) {
	password := "123456"
	hash := Hash(password)
	compare := Compare(password, hash)
	require.Equal(t, true, compare)
}

func TestBCrypt(t *testing.T) {
	password := "123"
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), 2)
	_ = err
	fmt.Println(bcrypt.CompareHashAndPassword(fromPassword, []byte(password)))
}

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
	testPwd := "123456"
	hash := Hash(testPwd)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(testPwd))
	if err != nil {
		t.Errorf("hash compare error: %s", err)
	}
}
func TestCompare(t *testing.T) {
	testPwd := "123456"
	hash := Hash(testPwd)
	if Compare(testPwd, hash) != true {
		t.Errorf("hash compare error: %s", "password not match")
	}
}

func TestGenHashedPassAndSalt(t *testing.T) {
	testPwd := "123456"
	hash := Hash(testPwd)
	compare := Compare(testPwd, hash)
	require.Equal(t, true, compare)
}

func TestBCrypt(t *testing.T) {
	testPwd := "123"
	bytes, err := bcrypt.GenerateFromPassword([]byte(testPwd), 2)
	_ = err
	fmt.Println(bcrypt.CompareHashAndPassword(bytes, []byte(testPwd)))
}

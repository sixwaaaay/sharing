package auth

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenToken(t *testing.T) {
	signer := NewJWTSigner("test")
	token, err := signer.GenerateToken(125, 60*60*24*7)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	t.Log(token)
}

func TestJWTSignerValidateToken(t *testing.T) {
	signer := NewJWTSigner("test")
	token, err := signer.GenerateToken(125, 60*60*24*7)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	t.Log(token)
	claims, err := signer.ValidateToken(token)
	require.NoError(t, err)
	require.NotNil(t, claims)
	require.Equal(t, int64(125), claims)
	t.Logf("%+v", claims)
}

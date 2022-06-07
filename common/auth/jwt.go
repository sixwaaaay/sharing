package auth

import (
	"bytelite/common/errorx"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type customClaims struct {
	UserID    int64 `json:"user_id"`
	ExpiredAt int64 `json:"expired_at"`
}

func NewCustomClaims(userID int64, expiredTime int64) *customClaims {
	return &customClaims{
		UserID:    userID,
		ExpiredAt: time.Now().Unix() + expiredTime,
	}
}

func (c *customClaims) Valid() error {
	if time.Now().Unix() > c.ExpiredAt {
		return errorx.NewDefaultError("token expired")
	}
	return nil
}

type JWTSigner struct {
	hmacSampleSecret []byte
}

func NewJWTSigner(secret string) *JWTSigner {
	return &JWTSigner{
		hmacSampleSecret: []byte(secret),
	}
}

// GenerateToken 生成token，过期时间以秒为单位
func (j *JWTSigner) GenerateToken(userID int64, expiryTime int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewCustomClaims(userID, expiryTime))
	return token.SignedString(j.hmacSampleSecret)
}

// ValidateToken 校验token，如果校验成功，返回user_id
func (j *JWTSigner) ValidateToken(tokenString string) (int64, error) {
	// validate the token whether is valid or not
	// if valid, return the user_id
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.hmacSampleSecret, nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
		return claims.UserID, nil
	}
	return 0, errorx.NewDefaultError("invalid token")
}

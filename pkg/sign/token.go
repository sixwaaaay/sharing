/*
 * Copyright (c) 2023 sixwaaaay.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sign

import (
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

type JWT struct {
	Secret string `yaml:"secret"` // JwtSecret is the secret used to sign the JWT
	TTL    int64  `yaml:"ttl"`    // JwtTTL is the time to live for the JWT in seconds
}

// customClaims are custom claims extending default ones.
type customClaims struct {
	Name   string `json:"name"`
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type SignOption struct {
	Username string
	UserID   int64
	Duration time.Duration
	Secret   []byte
}

func GenSignedToken(option SignOption) (string, error) {
	claims := &customClaims{
		UserID: strconv.FormatInt(option.UserID, 10),
		Name:   option.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(option.Duration)),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token
	return token.SignedString(option.Secret)
}

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
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func Middleware(secret []byte, requireToken bool) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		// 将 http.HandlerFunc 转换为 http.Handler
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				// If a token is not required, just pass through
				if !requireToken {
					r = withID(r, "0")
					next.ServeHTTP(w, r)
					return
				}
				// Otherwise, return an error
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
			// Parse the JWT token
			token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
				// Use your secret key or public key here to verify the signature
				return secret, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "Invalid Authorization Token")
				return
			}

			// Check if the token is valid
			if _, ok := token.Claims.(*customClaims); !ok || !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "Invalid Authorization Token")
				return
			}

			// Get the subject ID from the token
			id := token.Claims.(*customClaims).UserID

			// Pass the subject ID to the next handler
			r = withID(r, id)

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}

func withID(r *http.Request, id string) *http.Request {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "x-id", id)
	r = r.WithContext(ctx)
	return r
}

/*
 * Copyright (c) 2023-2024 sixwaaaay.
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

package api

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sixwaaaay/token"
)

// sign is a function that generates a JWT (JSON Web Token) with specific claims.
// It takes a secret key, a duration, an id, and a name as parameters.
// The secret key is used to sign the JWT.
// The duration is added to the current time to set the expiration time of the JWT.
// The id and name are added as claims to the JWT.
// It returns the signed JWT as a string and an error if any occurred during the signing process.
func sign(secret []byte, d time.Duration, id int64, name string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp":           time.Now().Add(d).Unix(), // seconds
			"iss":           "sharing",
			token.ClaimID:   strconv.FormatInt(id, 10),
			token.ClaimName: name,
		})
	s, err := t.SignedString(secret)
	return s, err
}

// Signer is a type alias for a function that takes an id and a name as parameters and returns a JWT and an error.
type Signer = func(id int64, name string) (string, error)

// SignFunc is a function that returns a Signer.
// It takes a secret key and a duration as parameters.
// The secret key is used to sign the JWT.
// The duration is added to the current time to set the expiration time of the JWT.
func SignFunc(secret []byte, d time.Duration) Signer {
	return func(id int64, name string) (string, error) {
		return sign(secret, d, id, name)
	}
}

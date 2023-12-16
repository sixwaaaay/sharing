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

package main

import (
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"net/http"
	"time"

	pb "codeberg.org/sixwaaaay/sharing-pb"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type Oauth2 struct {
	conf *oauth2.Config
	uc   pb.UserServiceClient
	signer
}

func NewOauth2(conf *oauth2.Config, uc pb.UserServiceClient, signer signer) *Oauth2 {
	return &Oauth2{conf: conf, uc: uc, signer: signer}
}

func (O *Oauth2) Update(e *echo.Echo) {
	e.Any("/oauth/github", O.Login)
	//	callback
	e.Any("/oauth/callback/github", O.Callback)
}

func (O *Oauth2) Login(c echo.Context) error {
	ref := c.Request().Header.Get("Referer") // get "Referer" from header
	if ref == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid referer")
	}
	cookie := http.Cookie{Name: "referer", Value: ref}
	c.SetCookie(&cookie)
	// redirect to GitHub oauth page
	url := O.conf.AuthCodeURL("", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusSeeOther, url)
}

type GithubUser struct {
	Name     string `json:"login"`
	AvtarUrl string `json:"avatar_url"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

func (O *Oauth2) Callback(c echo.Context) error {
	code := c.QueryParam("code")
	token, err := O.conf.Exchange(c.Request().Context(), code)
	if err != nil {
		return err
	}
	// get "referer" from cookie
	cookie, err := c.Cookie("referer")
	if err != nil {
		return err
	}
	user, err := O.getUserDetail(token)
	if err != nil {
		return err
	}
	reply, err := O.uc.Register(c.Request().Context(), &pb.RegisterRequest{
		Name:     user.Name,
		Email:    user.Email,
		Password: generateRandomString(16),
	})
	if err != nil {
		return err
	}
	// save user info to database
	_, err = O.uc.UpdateUser(c.Request().Context(), &pb.UpdateUserRequest{
		UserId:    reply.User.Id,
		Bio:       user.Bio,
		AvatarUrl: user.AvtarUrl,
	})
	if err != nil {
		return err
	}

	signedToken, err := O.signer(reply.User.Id, user.Name)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, cookie.Value+"?token="+signedToken)
}

type GithubEmail struct {
	Email      string `json:"email"`
	Primary    bool   `json:"primary"`
	Verified   bool   `json:"verified"`
	Visibility string `json:"visibility"`
}

func findPrimaryEmail(emails []GithubEmail) (string, error) {
	for _, email := range emails {
		if email.Primary {
			return email.Email, nil
		}
	}
	return "", errors.New("no primary email")
}

// getUserDetail is a method of the Oauth2 struct.
// It takes an oauth2.Token as an argument and returns a GithubUser struct and an error.
// This method is used to get the details of a GitHub user using the provided OAuth2 token.
// It makes two separate requests to the GitHub API:
// 1. To get the user's basic information (username, avatar URL, bio).
// 2. To get the user's email addresses.
// The method then finds the primary email address from the list of email addresses returned by the API.
// The user's basic information and primary email address are stored in a GithubUser struct and returned.
func (O *Oauth2) getUserDetail(token *oauth2.Token) (user GithubUser, err error) {
	// Create a new HTTP client
	client := http.Client{}

	// Make a GET request to the GitHub API to get the user's basic information
	body, err := getFromBody(token.AccessToken, client, "https://api.github.com/user")
	if err != nil {
		return user, err
	}
	defer body.Close()

	// Decode the response body into the GithubUser struct
	if err := json.NewDecoder(body).Decode(&user); err != nil {
		return user, err
	}

	// Make a GET request to the GitHub API to get the user's email addresses
	var emails []GithubEmail
	body, err = getFromBody(token.AccessToken, client, "https://api.github.com/user/emails")
	if err != nil {
		return user, err
	}
	defer body.Close()

	// Decode the response body into a slice of GithubEmail structs
	if err := json.NewDecoder(body).Decode(&emails); err != nil {
		return user, err
	}

	// Find the primary email address from the list of email addresses
	email, err := findPrimaryEmail(emails)
	if err != nil {
		return user, err
	}

	// Store the primary email address in the GithubUser struct
	user.Email = email

	// Return the GithubUser struct and any error that occurred
	return user, nil
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012345"

// generate a random string with given length
func generateRandomString(length int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		randomString[i] = charset[rand.Intn(len(charset))]
	}
	return string(randomString)
}

func getFromBody(token string, client http.Client, url string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

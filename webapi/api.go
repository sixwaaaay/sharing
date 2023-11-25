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
	"net/http"

	"github.com/sixwaaaay/shauser/internal/logic"
	"github.com/sixwaaaay/shauser/user"
)

type WebApi struct {
	GetUserLogic  *logic.GetUserLogic
	GetUsersLogic *logic.GetUsersLogic
}

func NewWebApi(getUserLogic *logic.GetUserLogic, getUsersLogic *logic.GetUsersLogic) *WebApi {
	return &WebApi{GetUserLogic: getUserLogic, GetUsersLogic: getUsersLogic}
}

func (api *WebApi) GetUserHandler(w http.ResponseWriter, r *http.Request) error {
	var in user.GetUserRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		return err
	}
	reply, err := api.GetUserLogic.GetUser(r.Context(), &in)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(reply)
}

func (api *WebApi) GetUsersHandler(w http.ResponseWriter, r *http.Request) error {
	var in user.GetUsersRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		return err
	}
	reply, err := api.GetUsersLogic.GetUsers(r.Context(), &in)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(reply)
}

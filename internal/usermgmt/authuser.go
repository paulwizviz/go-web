// Copyright 2020 Paul Sitoh
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package usermgmt

import (
	"goweb/internal"
	"net/http"

	"github.com/gorilla/mux"
)

// AuthUser represents a user who has been verified to use this app
type AuthUser struct {
	Authenticator func(id string, secrets string) ([]byte, error)
}

func (a *AuthUser) Handler(rw http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	if req.URL.Path != internal.URLAuthPath {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Content-Type", "application/json")

	param := mux.Vars(req)
	id := param["id"]
	secrets := param["secrets"]

	token, err := a.Authenticator(id, secrets)
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(token)
}

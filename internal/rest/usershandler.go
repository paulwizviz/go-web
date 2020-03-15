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

package rest

import (
	"encoding/json"
	"net/http"
)

type user struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func mockUsers() []user {
	return []user{
		user{
			Name:     "Albert",
			Password: "albert1234",
		},
		user{
			Name:     "Beatrice",
			Password: "beatrice1234",
		},
	}
}

func usersPayload() []byte {
	payload, _ := json.Marshal(mockUsers())
	return payload
}

func usershandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))
	w.Write(usersPayload())
}

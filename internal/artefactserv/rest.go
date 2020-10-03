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

package server

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func authHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	if req.URL.Path != URLAuthPath {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}
	defer req.Body.Close()

	rw.Header().Set(HTTPHeaderAccessControllerAllowOrigin, "*")
	rw.Header().Set(HTTPHeaderContentType, "application/json")

	rw.WriteHeader(http.StatusOK)
	rw.Write(body)

}

func RESTRun(router *mux.Router) {
	router.HandleFunc(URLAuthPath, authHandler)
	router.Use(mux.CORSMethodMiddleware(router))
}

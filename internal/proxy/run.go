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

package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

// Handler responsible for routing path / to web server
func handleWebRequest(res http.ResponseWriter, req *http.Request) {

	port := "3000"
	webURL, err := url.ParseRequestURI(fmt.Sprintf("http://localhost:%v/", port))
	if err != nil {
		log.Printf("Web server error %v", err)
	}

	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.URL = webURL

	proxy := httputil.NewSingleHostReverseProxy(webURL)
	proxy.ServeHTTP(res, req)
}

// Handler responsible for routing path /api to REST server
func handleAPIRequest(res http.ResponseWriter, req *http.Request) {

	port := "9000"
	apiURL, err := url.ParseRequestURI(fmt.Sprintf("http://localhost:%v/", port))
	if err != nil {
		log.Printf("Web server error %v", err)
	}

	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.URL = apiURL

	proxy := httputil.NewSingleHostReverseProxy(apiURL)
	proxy.ServeHTTP(res, req)

}

func Run() {
	port := "80"
	router := mux.NewRouter()

	router.HandleFunc("/", handleWebRequest)
	router.HandleFunc("/api", handleAPIRequest)

	log.Printf("Proxy starting on %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), router))
}

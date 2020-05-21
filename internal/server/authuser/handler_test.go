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

package authuser

import (
	"encoding/json"
	"goweb/internal"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvalidAuthHandlerMethod(t *testing.T) {
	req, err := http.NewRequest("GET", internal.URLAuthPath, nil)
	if err != nil {
		t.Fatalf("Unable to create a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(Handler).ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("Expected status: %v Got status: %v", http.StatusUnauthorized, rr.Code)
	}
}

func TestInvalidAuthHandlerPath(t *testing.T) {
	req, err := http.NewRequest("POST", "/users", nil)
	if err != nil {
		t.Fatalf("Unable to create a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(Handler).ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("Expected status: %v Got status: %v", http.StatusUnauthorized, rr.Code)
	}
}

func mockAuthUserRequestBody(t *testing.T) []byte {

	reqBody := struct {
		ID      string `json:"id"`
		Secrets string `json:"secrets"`
	}{
		ID:      "John",
		Secrets: "1we5d",
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Unable to mock request body %v", err)
	}

	return bodyBytes
}

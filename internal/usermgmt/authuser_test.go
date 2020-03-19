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
	"bytes"
	"encoding/json"
	"goweb/internal"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockAuthUserRequestBody(t *testing.T) []byte {

	reqBody := struct {
		ID      string `json:"id"`
		Secrets string `json:"secrets"`
	}{
		ID:      "Paul",
		Secrets: "1%we5d",
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Unable to mock request body %v", err)
	}

	return bodyBytes
}

func TestAuthUserInvalidMethod(t *testing.T) {

	expectedToken := []byte("1fsd23")

	authUser := AuthUser{}
	authUser.Authenticator = func(id string, secrets string) ([]byte, error) {
		return expectedToken, nil
	}

	bodyBytes := mockAuthUserRequestBody(t)
	bytesReader := bytes.NewReader(bodyBytes)

	req, err := http.NewRequest("GET", internal.URLAuthPath, bytesReader)
	if err != nil {
		t.Fatalf("Unable to create a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(authUser.Handler).ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("Expected status: %v Got status: %v", http.StatusUnauthorized, rr.Code)
	}
}

func TestAuthUserInvalidPath(t *testing.T) {

	expectedToken := []byte("1fsd23")

	authUser := AuthUser{}
	authUser.Authenticator = func(id string, secrets string) ([]byte, error) {
		return expectedToken, nil
	}

	bodyBytes := mockAuthUserRequestBody(t)
	bytesReader := bytes.NewReader(bodyBytes)

	req, err := http.NewRequest("POST", "/users", bytesReader)
	if err != nil {
		t.Fatalf("Unable to create a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(authUser.Handler).ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("Expected status: %v Got status: %v", http.StatusUnauthorized, rr.Code)
	}
}

func TestAuthUserPayload(t *testing.T) {
	expectedToken := []byte("1fsd23")

	authUser := AuthUser{}
	authUser.Authenticator = func(id string, secrets string) ([]byte, error) {
		return expectedToken, nil
	}

	bodyBytes := mockAuthUserRequestBody(t)
	bytesReader := bytes.NewReader(bodyBytes)

	req, err := http.NewRequest("POST", internal.URLAuthPath, bytesReader)
	if err != nil {
		t.Fatalf("Unable to create a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(authUser.Handler).ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Expected status: %v Got status: %v", http.StatusUnauthorized, rr.Code)
	}

	resBody, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("Unable to read response body: %v", err)
	}

	if bytes.Compare(resBody, expectedToken) != 0 {
		t.Fatalf("Expected: %v Got: %v", string(expectedToken), string(resBody))
	}
}

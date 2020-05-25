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

package file

import (
	"goweb/internal/authuser"
)

type mockCredentialRepok struct{}

func (c *mockCredentialRepok) FindLoginCred(id string) (*authuser.LoginCredential, error) {
	return &authuser.LoginCredential{
		ID:      "id",
		Secrets: "secrets",
	}, nil
}

func (c *mockCredentialRepok) FindAccessCred(id string) (*authuser.AccessCredential, error) {
	return &authuser.AccessCredential{
		ID:          "Test",
		DisplayName: "Test",
		AccessToken: "Test",
		AccessRole:  "Test",
	}, nil
}

func NewMockCredentialRepoService() authuser.CredentialRepo {
	return &mockCredentialRepok{}
}

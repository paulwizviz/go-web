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
	"fmt"
	"reflect"
)

const errorPrefix = "Missing:"

type fieldError struct {
	FieldName string
}

func (e *fieldError) Error() string {
	return fmt.Sprintf("%v %v", errorPrefix, e.FieldName)
}

// EmptyFieldError identify erroneous field
func EmptyFieldError(err error) string {
	return err.(*fieldError).FieldName
}

func isEmptyString(value string) bool {
	if len(value) == 0 {
		return true
	}
	return false
}

// LoginCredential associated with a person presented for authentication
type LoginCredential struct {
	// ID assigned to person holding this credential
	ID string `json:"id"`
	// Secrets provided by the person with this credential
	Secrets string `json:"secrets"`
}

// VerifyLoginCredForm determines if all fields have been set
func VerifyLoginCredForm(cred LoginCredential) error {

	field := reflect.Indirect(reflect.ValueOf(cred))
	if isEmptyString(cred.ID) {
		return &fieldError{
			FieldName: field.Type().Field(0).Name,
		}
	}
	if isEmptyString(cred.Secrets) {
		return &fieldError{
			FieldName: field.Type().Field(1).Name,
		}
	}
	return nil
}

// AccessCredential contains information associated with a person
type AccessCredential struct {
	// ID is a unique string value associated with a user.
	ID string `json:"id"`
	// DisplayName is an optional representation of a user name
	// that can be read by a person
	DisplayName string `json:"display_name"`
	// JWTToken is an optional representation of a JWT Token.
	AccessToken string `json:"access_token"`
	// AccessRole represents role
	AccessRole string `json:"access_role"`
}

// VerifyAccessCredForm determines if access credential has at least ID field set
func VerifyAccessCredForm(cred AccessCredential) error {
	if isEmptyString(cred.ID) {
		return &fieldError{
			FieldName: "ID",
		}
	}
	return nil
}

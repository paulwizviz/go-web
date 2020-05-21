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
	"testing"
)

func TestVerifyLoginCredForm(t *testing.T) {

	dataSet := []struct {
		cred     LoginCredential
		expected error
	}{
		{
			cred: LoginCredential{},
			expected: &fieldError{
				FieldName: "ID",
			},
		},
		{
			cred: LoginCredential{
				ID: "id",
			},
			expected: &fieldError{
				FieldName: "Secrets",
			},
		},
	}

	for _, dataItem := range dataSet {
		err := VerifyLoginCredForm(dataItem.cred)
		if err.Error() != dataItem.expected.Error() {
			t.Fatalf("Expected: %v Got: %v", dataItem.expected, err)
		}
	}
}

func TestVerifyAccessCredForm(t *testing.T) {
	dataSet := []struct {
		input    AccessCredential
		expected error
	}{
		{
			input: AccessCredential{},
			expected: &fieldError{
				FieldName: "ID",
			},
		},
	}

	for _, dataItem := range dataSet {
		err := VerifyAccessCredForm(dataItem.input)
		if err.Error() != dataItem.expected.Error() {
			t.Fatalf("Expected: %v Got: %v", dataItem.expected, err)
		}
	}
}

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

package cli

import "testing"

func TestRestEndCmdName(t *testing.T) {
	builder := RESTEndCmdBuilder{}
	builder.services = func() {}
	cmd := builder.cli()
	expected := "rest"
	got := cmd.Use
	if expected != got {
		t.Errorf("Expected: %v Got: %v", expected, got)
	}
}

func TestRestEndCmdFlag(t *testing.T) {
	builder := RESTEndCmdBuilder{}
	builder.services = func() {
		expected := 8080
		got := builder.port
		if expected != got {
			t.Errorf("Expected: %v Got: %v", expected, got)
		}
	}
	cmd := builder.cli()
	cmd.Flags().IntVarP(&builder.port, "port", "p", 8080, "default port 8080")
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Unable to execute command. Reason: %v", err)
	}
}

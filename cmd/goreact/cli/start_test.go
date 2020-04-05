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

func TestStartCmdName(t *testing.T) {

	builder := StartCmdBuilder{}
	builder.service = func() {
	}

	cmd := builder.cli()
	expected := "start"
	got := cmd.Use
	if expected != got {
		t.Errorf("Expected: %v Got: %v", expected, got)
	}
}

func TestStartCmdNoUIFlag(t *testing.T) {
	builder := StartCmdBuilder{
		ui: true,
	}
	builder.service = func() {
		expected := false
		got := builder.ui
		if expected != got {
			t.Errorf("Expected: %v Got: %v", expected, got)
		}
	}
	cmd := builder.cli()
	cmd.Flags().BoolVar(&builder.ui, "ui", false, "default true")
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Unable to execute command. Reason: %v", err)
	}
}

func TestStartCmdPortFlag(t *testing.T) {
	builder := StartCmdBuilder{
		port: 8080,
	}
	builder.service = func() {
		expected := 80
		got := builder.port
		if expected != got {
			t.Errorf("Expected: %v Got: %v", expected, got)
		}
	}
	cmd := builder.cli()
	cmd.Flags().IntVarP(&builder.port, "port", "p", 80, "startup default port 80")
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Unable to execute command. Reason: %v", err)
	}
}

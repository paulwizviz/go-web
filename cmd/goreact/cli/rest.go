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

import (
	"fmt"
	"goweb/internal/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

type RESTEndCmdBuilder struct {
	port     int
	services func()
}

func (s *RESTEndCmdBuilder) cli() *cobra.Command {
	return &cobra.Command{
		Use:   "rest",
		Short: "starts the orchestrator RESTFul server only.",
		Run: func(cmd *cobra.Command, args []string) {
			s.services()
		},
	}
}

var restendCmdBuilder = RESTEndCmdBuilder{}

func init() {
	restendCmdBuilder.services = func() {
		router := mux.NewRouter()
		server.RESTRun(router)
		log.Printf("Starting on port %v", restendCmdBuilder.port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", frontendCmdBuilder.port), router))
	}
}

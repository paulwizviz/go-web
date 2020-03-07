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
	"go-web/internal/rest"
	"go-web/internal/webserver"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

func goreactUseCase() string {
	return `go-react is a example cli toolkit to startup a ReactJS web`
}

var rootCmd = &cobra.Command{
	Use:   "go-react",
	Short: "go-react is a cli app",
	Long:  goreactUseCase(),
	Run: func(cmd *cobra.Command, args []string) {
		port := "80"
		router := mux.NewRouter()
		rest.Run(router)
		webserver.Run(router)
		log.Printf("Starting on port %v", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), router))
	},
}

// Execute is the cli entry point
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

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
	"log"
	"net/http"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/paulwizviz/go-react/internal/rest"
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
		go func() {
			rest.Start()
		}()
		router := mux.NewRouter()
		router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("../../../web").HTTPBox()))
		log.Printf("Starting react %d", 3000)
		log.Fatal(http.ListenAndServe("0.0.0.0:3000", router))
	},
}

// Execute is the cli entry point
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

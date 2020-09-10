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

package main

import (
	"flag"
	"fmt"
	"goweb/internal/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const defaultPort = 80
const devPort = 9000

func main() {

	devPtr := flag.Bool("noui", false, "Run with no ui")

	router := mux.NewRouter()

	if *devPtr == true {
		server.RESTRun(router)
		log.Printf("Starting REST server on port %v", devPort)
		log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", devPort), router))
	} else {
		server.RESTRun(router)
		server.RunWeb(router)
		log.Printf("Starting with UI on port %v", defaultPort)
		log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", defaultPort), router))
	}

}

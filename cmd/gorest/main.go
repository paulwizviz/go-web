package main

import (
	"fmt"
	server "goweb/internal/artefactserv"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = 9000

func main() {
	router := mux.NewRouter()
	server.RESTRun(router)
	log.Printf("Starting REST server on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), router))
}

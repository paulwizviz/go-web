package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	router := mux.NewRouter()
	router.HandleFunc("/", homehandler)
	router.HandleFunc("/users", usershandler)
	router.Use(mux.CORSMethodMiddleware(router))
	log.Printf("Starting rest server %v", 9000)
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", router))
}

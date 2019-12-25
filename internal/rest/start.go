package rest

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/", homehandler)
	r.HandleFunc("/users", usershandler)
	r.Use(mux.CORSMethodMiddleware(r))
	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:9000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Starting rest server %v", srv.Addr)
	srv.ListenAndServe()
}

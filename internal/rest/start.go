package rest

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/", homehandler)
	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:9000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}

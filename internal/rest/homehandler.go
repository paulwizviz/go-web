package rest

import (
	"fmt"
	"net/http"
)

func homehandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader((http.StatusOK))
	fmt.Fprintf(w, "Hello")
}

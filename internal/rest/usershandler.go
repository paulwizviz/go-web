package rest

import (
	"fmt"
	"net/http"
)

func usershandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))
	fmt.Fprintf(w, `{users: [{"name":"Albert"}.{"name":"Beatrice"}]}`)
}

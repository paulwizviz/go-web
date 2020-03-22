package server

import (
	"goweb/internal"
	"goweb/internal/usermgmt"
	"net/http"

	"github.com/gorilla/mux"
)

func authUserHandler(rw http.ResponseWriter, req *http.Request) {
	authUser := usermgmt.AuthUser{}
	authUser.Authenticator = func(id string, secret string) ([]byte, error) {
		return []byte("bearer: 12d4f122"), nil
	}
	authUser.Handler(rw, req)
}

func RESTRun(router *mux.Router) {
	router.HandleFunc(internal.URLAuthPath, authUserHandler)
	router.Use(mux.CORSMethodMiddleware(router))
}

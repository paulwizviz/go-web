package cli

import (
	"fmt"
	"goweb/internal/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

type FrontendCmdBuilder struct {
	port     int
	services func()
}

func (f *FrontendCmdBuilder) cli() *cobra.Command {
	return &cobra.Command{
		Use:   "frontend",
		Short: "starts the UI and RESTFul servers.",
		Run: func(cmd *cobra.Command, args []string) {
			f.services()
		},
	}
}

var frontendCmdBuilder = FrontendCmdBuilder{}

func init() {
	frontendCmdBuilder.services = func() {
		router := mux.NewRouter()
		server.RESTRun(router)
		server.WebRun(router)
		log.Printf("Starting on port %v", frontendCmdBuilder.port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", frontendCmdBuilder.port), router))
	}
}

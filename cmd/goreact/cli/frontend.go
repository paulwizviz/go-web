package cli

import (
	"fmt"
	"goweb/internal/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var frontendCmd = &cobra.Command{
	Use:   "frontend",
	Short: "frontend is a subcommand start frontend services",
	Run: func(cmd *cobra.Command, args []string) {
		router := mux.NewRouter()
		server.RESTRun(router)
		server.WebRun(router)
		log.Printf("Starting on port %v", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), router))
	},
}

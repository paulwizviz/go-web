package cli

import (
	"fmt"
	"go-web/internal/rest"
	"go-web/internal/webserver"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var frontendCmd = &cobra.Command{
	Use:   "frontend",
	Short: "frontend is a subcommand start frontend services",
	Run: func(cmd *cobra.Command, args []string) {
		port := "80"
		router := mux.NewRouter()
		rest.Run(router)
		webserver.Run(router)
		log.Printf("Starting on port %v", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), router))
	},
}

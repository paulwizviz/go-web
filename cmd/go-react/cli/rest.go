package cli

import (
	"fmt"
	"go-web/internal/rest"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "Subcommand to start rest server only",
	Run: func(cmd *cobra.Command, args []string) {
		port := "9000"
		router := mux.NewRouter()
		rest.Run(router)
		log.Printf("Starting on port %v", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), router))
	},
}

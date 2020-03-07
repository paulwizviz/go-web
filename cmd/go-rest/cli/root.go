package cli

import (
	"fmt"
	"go-web/internal/rest"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var port int

func appName() string {
	return "go-rest"
}

func useCase() string {
	return fmt.Sprintf(`%v is a RESTFul server`, appName())
}

var rootCmd = &cobra.Command{
	Use:   appName(),
	Short: fmt.Sprintf("%v is a RESTFul server", appName()),
	Run: func(cmd *cobra.Command, args []string) {
		router := mux.NewRouter()
		rest.Run(router)
		log.Printf("Starting rest server on port %v", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), router))
	},
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 80, "startup default port 80")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

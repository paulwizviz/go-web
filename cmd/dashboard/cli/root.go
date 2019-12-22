package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func goreactUseCase() string {
	return `go-react is a example cli toolkit to startup a ReactJS web`
}

var rootCmd = &cobra.Command{
	Use:   "go-react",
	Short: "go-react is a cli app",
	Long:  goreactUseCase(),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello world")
	},
}

// Execute is the cli entry point
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

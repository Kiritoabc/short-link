package app

import (
	"github.com/Kiritoabc/short-link/backend/pkg/config"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := cobra.Command{
		Use:  "backend-server",
		Long: "Short Link Service Backend",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
	// apply config
	applyConfig(&cmd)
	return &cmd
}

func applyConfig(cmd *cobra.Command) {
	flags := config.Flags
	for _, flag := range flags {
		cmd.Flags().StringVar(&flag.Value, flag.Name, flag.Value, flag.Description)
	}
}

// init
func run() {

}

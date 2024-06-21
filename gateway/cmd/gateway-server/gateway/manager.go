package gateway

import (
	"github.com/Kiritoabc/short-link/gateway/cmd/gateway-server/pkg/config"
	"github.com/Kiritoabc/short-link/gateway/cmd/gateway-server/pkg/midware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := cobra.Command{
		Use:  "gateway",
		Long: "gateway server",
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
		cmd.Flags().StringVarP(&flag.Value, flag.Name, flag.Name, flag.Value, flag.Description)
	}
}

func run() {
	engine := gin.Default()
	engine.Use(midware.LoadBalancerMiddleware)
	if err := engine.Run(config.Port.Value); err != nil {
		panic(err)
	}
}

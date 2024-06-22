package app

import (
	"github.com/Kiritoabc/short-link/backend/pkg/config"
	"github.com/Kiritoabc/short-link/backend/pkg/database"
	router2 "github.com/Kiritoabc/short-link/backend/pkg/router"
	"github.com/Kiritoabc/short-link/backend/pkg/utils/snowflake"
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
		//cmd.Flags().StringP(flag.Name, "", flag.Value, flag.Description)
	}
}

// init
func run() {
	// 舒适化数据库连接
	err := database.Init()
	if err != nil {
		panic(err)
	}
	// init snowflake
	snowflake.InitSnowFlake()

	// gin router
	router := router2.InitGinRouter()
	if err := router.Run(config.Port.Value); err != nil {
		panic(err)
	}
}

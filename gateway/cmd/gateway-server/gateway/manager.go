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
	applyLoadBalancer()
	return &cmd
}

func applyConfig(cmd *cobra.Command) {
	flags := config.Flags
	for _, flag := range flags {
		cmd.Flags().StringVarP(&flag.Value, flag.Name, flag.Name, flag.Value, flag.Description)
	}
}

// 选项模式？

func applyLoadBalancer() {
	// 测试
	// 初始化，负载均衡的IP地址
	midware.Ip = append(midware.Ip, "http://127.0.0.1:8081", "http://127.0.0.1:8081")
	// load balance
	// 1. random
	// 2. round robin
	// 3. least connection
	// 4. least time
	// 5. least time with least connection
	// 6. least time with least connection with least time
	// 7. least time with least connection with least time with least connection
	// 8. least time with least connection with least time with least connection with least time
	// 9. least time with least connection with
}

func run() {
	engine := gin.Default()
	engine.Use(midware.LoadBalancerMiddleware)
	if err := engine.Run(config.Port.Value); err != nil {
		panic(err)
	}
}
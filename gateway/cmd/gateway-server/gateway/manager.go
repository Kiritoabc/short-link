package gateway

import (
	"fmt"
	"github.com/Kiritoabc/short-link/gateway/cmd/pkg/config"
	"github.com/Kiritoabc/short-link/gateway/cmd/pkg/midware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"sort"
)

func Command() *cobra.Command {
	cmd := cobra.Command{
		Use:   "gateway",
		Short: "gateway server",
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
		cmd.Flags().StringVar(&flag.Value, flag.Name, flag.Value, flag.Description)
	}
}

// 选项模式？

func applyLoadBalancer() {
	// 测试
	// 初始化，负载均衡的IP地址
	midware.Ip = append(midware.Ip, "http://127.0.0.1:8081", "http://127.0.0.1:8082")
	// 测试
	for i, v := range midware.Ip {
		midware.WeightIps = append(midware.WeightIps, midware.WeightNode{Node: v, Weight: i})
	}
	// sort ? 不一定需要
	sort.Slice(midware.WeightIps, func(i, j int) bool {
		return midware.WeightIps[i].Weight > midware.WeightIps[j].Weight
	})

	// 初始化 负载均衡算法
	switch config.ProxyModel.Value {
	case "pool":
		midware.InitPoolServer()
	case "rand":
		midware.InitRandServer()
	case "randAndPool":
		midware.InitRandPoolServer()
	case "weightPool":
		midware.InitWeightPoolServer()
	case "hash":
		midware.InitHashRing()
	default:
		fmt.Println("proxy model error")
		midware.InitPoolServer()
	}
}

func run() {
	engine := gin.Default()
	engine.Use(midware.LoadBalancerMiddleware)
	if err := engine.Run(config.Port.Value); err != nil {
		panic(err)
	}
}

package midware

import (
	"github.com/Kiritoabc/short-link/gateway/cmd/gateway-server/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/url"
)

var Ip = []string{}

var poolIndex = 0

// 初始化我们需要的IP

// LoadBalancerMiddleware 负载均衡的中间件
func LoadBalancerMiddleware(ctx *gin.Context) {
	// 开启代理==> 不是开启代理，而是转发
	poolLoadBalance(ctx)
	return
}

// 轮询

func poolLoadBalance(ctx *gin.Context) {
	n := len(Ip)
	if poolIndex >= n {
		poolIndex = 0
	}
	// 拿到代理的Ip地址
	proxyIp := Ip[poolIndex]
	target, err := url.Parse(proxyIp)
	poolIndex++
	if err != nil {
		panic(err)
	}
	// 不是开启代理 ==> 而是转发
	body := utils.ForwardRequest(target, ctx.Request)
	ctx.Writer.Write(body)
}

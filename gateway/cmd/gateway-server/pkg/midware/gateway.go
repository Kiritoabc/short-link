package midware

import (
	"github.com/Kiritoabc/short-link/gateway/cmd/gateway-server/pkg/config"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"net/url"
)

// LoadBalancerMiddleware 负载均衡的中间件
func LoadBalancerMiddleware(ctx *gin.Context) {
	// 这里轮训获取代理地址
	proxyUrl := config.ProxyPort.Value
	target, err := url.Parse(proxyUrl)
	if err != nil {
		panic(err)
	}
	// 开启代理
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
	ctx.Next()
}

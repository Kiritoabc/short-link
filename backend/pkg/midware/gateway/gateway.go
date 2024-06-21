package gateway

import (
	"github.com/gin-gonic/gin"
	"math/rand"
)

func LoadBalancerMiddleware() gin.HandlerFunc {
	// 服务器列表
	servers := []string{"http://localhost:8080", "http://localhost:8081", "http://localhost:8082"}

	return func(c *gin.Context) {
		// 选择服务器
		server := servers[rand.Intn(len(servers))]

		// 转发请求到选择的服务器
		c.Request.URL.Host = server
		c.Request.URL.Scheme = "http"

		// 调用下一个中间件或处理函数
		c.Next()
	}
}

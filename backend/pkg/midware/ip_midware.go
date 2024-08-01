package midware

import (
	"github.com/gin-gonic/gin"
)

// IpMidwareHandFunc ip 中间件
func IpMidwareHandFunc(ctx *gin.Context) {
	// 获取发送者的 IP 地址
	//logger.GetLogger().Info(ctx.RemoteIP())
	println(ctx.RemoteIP())
}

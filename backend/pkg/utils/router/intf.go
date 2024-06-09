package router

import "github.com/gin-gonic/gin"

type ApiController interface {
	GetGroup(engine *gin.Engine) *gin.RouterGroup
	RegisterRouter(group *gin.RouterGroup)
}

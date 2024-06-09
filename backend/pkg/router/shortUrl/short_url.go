package shortUrl

import (
	shortUrlHandler "github.com/Kiritoabc/short-link/backend/pkg/handler/shortUrl"
	"github.com/Kiritoabc/short-link/backend/pkg/utils"
	routerutil "github.com/Kiritoabc/short-link/backend/pkg/utils/router"
	"github.com/gin-gonic/gin"
)

func NewShortUrlRouter() routerutil.ApiController {
	return &shortUrlRouter{
		handler: shortUrlHandler.NewShortUrlHandler(),
	}
}

type shortUrlRouter struct {
	handler shortUrlHandler.ShortURLIntf
}

func (s *shortUrlRouter) GetGroup(engine *gin.Engine) *gin.RouterGroup {
	return engine.Group(utils.OpenApiGroup)
}

func (s *shortUrlRouter) RegisterRouter(group *gin.RouterGroup) {
	group.GET("/hello", s.Hello)
}

// Hello hello接口
//
//	@Tags			短链接健康检查
//	@Summary		健康检查
//	@Description	健康检查
//	@Router			/openapis/hello [get]
//	@Success		200	{object}	utils.Response{msg=string}
func (s *shortUrlRouter) Hello(ctx *gin.Context) {

	utils.OkWithMessage("hello world", ctx)
	return
}

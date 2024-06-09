package shortUrl

import (
	shortUrlHandler "github.com/Kiritoabc/short-link/backend/pkg/handler/shortUrl"
	"github.com/Kiritoabc/short-link/backend/pkg/utils"
	routerutil "github.com/Kiritoabc/short-link/backend/pkg/utils/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
	group.POST("/short", s.GetShortLink)
	group.GET("/short/*shortlink", s.GetLink)
}

var Hmap = make(map[string]string)

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

// GetShortLink 获取短链接
//
// @Tags			获取锻链接
// @Summary		获取锻链接
// @Description	获取锻链接
// @Router			/openapis/short 	[post]
// @Param			url					formData	string	true	"原始的长链接"
// @Success		200					{object}	utils.Response{msg=string}
func (s *shortUrlRouter) GetShortLink(ctx *gin.Context) {
	url, ok := ctx.GetPostForm("url")
	if !ok {
		utils.FailWithMessage("url参数错误", ctx)
		return
	}
	// 使用snowflake测试
	shortUrl, err := s.handler.GetShortLink(url)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	Hmap[shortUrl] = url
	utils.OkWithData(shortUrl, ctx)
}

// GetLink 获取短链接
//
// @Tags			短链接跳转
// @Summary		短链接跳转
// @Description	短链接跳转
// @Router			/openapis/ 				[get]
// @Param			shortlink	path		string	true	"短链接"
// @Success		200			{object}	utils.Response{msg=string}
func (s *shortUrlRouter) GetLink(ctx *gin.Context) {
	shortlink := ctx.Param("shortlink")
	split := strings.Split(shortlink, "/")
	url := Hmap[split[1]]
	if url == "" {
		utils.FailWithMessage("短链接不存在", ctx)
		return
	}
	// 将链接放入Location，使用302跳转,只跳转到url，不要包含当前的业务
	ctx.Redirect(http.StatusFound, url)
}

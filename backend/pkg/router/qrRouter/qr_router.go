package qrRouter

import (
	"github.com/Kiritoabc/short-link/backend/pkg/handler/qrHandler"
	"github.com/Kiritoabc/short-link/backend/pkg/utils"
	routerutil "github.com/Kiritoabc/short-link/backend/pkg/utils/router"
	"github.com/gin-gonic/gin"
)

func NewQrRouter() routerutil.ApiController {
	return &QrRouter{
		handler: qrHandler.NewQrHandler(),
	}
}

type QrRouter struct {
	handler qrHandler.QrIntf
}

func (q *QrRouter) GetGroup(engine *gin.Engine) *gin.RouterGroup {
	return engine.Group(utils.OpenApiGroup + "/qr")
}

func (q *QrRouter) RegisterRouter(group *gin.RouterGroup) {
	group.GET("/generate", q.GenerateQrCode)
}

func (q *QrRouter) GenerateQrCode(ctx *gin.Context) {
	url := ctx.Query("url")
	qrCode, err := q.handler.GenerateQrCode(url)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.OkWithQrCode(qrCode, ctx)
}

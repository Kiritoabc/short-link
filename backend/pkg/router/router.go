package router

import (
	"github.com/Kiritoabc/short-link/backend/pkg/config"
	"github.com/Kiritoabc/short-link/backend/pkg/router/shortUrl"
	routerutil "github.com/Kiritoabc/short-link/backend/pkg/utils/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strings"
)

var controllerFuncs = []func() routerutil.ApiController{
	shortUrl.NewShortUrlRouter,
}

// swag init -g router.go -o ../../../docs --parseDependency --parseInternal

// InitGinRouter init gin router
//
//	@title		Short Link API
//	@version	1.0.0
func InitGinRouter() *gin.Engine {

	// get gin default engine
	engine := gin.Default()
	engine.Use()

	// add mid ware handlers

	applySwagger(engine)

	// register router
	for _, cf := range controllerFuncs {
		c := cf()
		router := c.GetGroup(engine)
		c.RegisterRouter(router)
	}
	return engine
}

func applySwagger(engine *gin.Engine) {
	if strings.EqualFold(config.EnableSwagger.Value, "true") {
		// swagger
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

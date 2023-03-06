package initialize

import (
	"erp-web/controller"
	"erp-web/global"
	"erp-web/middlewares"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	g := gin.Default()
	routerGroup := g.Group("/pp/sfc/sfc102")
	{
		sfc102Controller := controller.NewSfc102Controller(global.DB)
		routerGroup.POST("/getOrdnoList", sfc102Controller.GetOrdnoList)
	}
	//配置跨域
	g.Use(middlewares.Cors())

	return g
}

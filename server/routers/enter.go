package routers

import (
	"github.com/gin-gonic/gin"
)

type RouterGroup struct{
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	// 设置开发显示模式
	// gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{ apiRouterGroup }
	routerGroupApp.SettingsRouter()
	return router
}

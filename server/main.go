package main

import (
	"server/core"
	"server/global"
	"server/routers"
)

func main() {
	// 读取配置文件
	core.InitConfig()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()
	// router
	addr := global.Config.System.Addr()
	global.Log.Infof("server运行在: %s",addr)
	router := routers.InitRouter()
	router.Run(addr)
}
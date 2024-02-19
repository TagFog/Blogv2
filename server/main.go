package main

import (
	"server/core"
	"server/global"
)

func main() {
	// 读取配置文件
	core.InitConfig()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()
}
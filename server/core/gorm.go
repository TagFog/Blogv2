package core

import (
	"fmt"
	"log"
	"server/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
// 数据库连接
func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Warnln("未配置MySQL,取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()
	var mysqlLogger logger.Interface
	// 开发环境
	if global.Config.System.Env == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	}else{
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil{
		log.Fatalf(fmt.Sprintf("[%s] mysql连接失败",dsn))
		panic(err)
	}
	sqlDB,_ := db.DB()
	sqlDB.SetMaxIdleConns(10) //最大空闲连接数
	sqlDB.SetMaxOpenConns(100) //最大容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) //最大连接复用时间
	return db
}

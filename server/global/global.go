package global

import (
	"server/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 用于保存配置文件
var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)

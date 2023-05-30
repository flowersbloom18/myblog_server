package global

import (
	"github.com/cc14514/go-geoip2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"myblog_server/config"
)

var (
	Config   *config.Config
	Log      *logrus.Logger
	MysqlLog logger.Interface
	DB       *gorm.DB
	AddrDB   *geoip2.DBReader
)

package global

import (
	"github.com/go-redis/redis"
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
	Redis    *redis.Client
)

package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	NewConfig *Config
	DB        *gorm.DB
	Log       *logrus.Logger
	Router    *gin.Engine
)

type Config struct {
	Mysql   Mysql   `yaml:"mysql"`
	Redis   Redis   `yaml:"Redis"`
	MongoDB MongoDB `yaml:"MongoDB"`
	ES      ES      `yaml:"ES"`
}

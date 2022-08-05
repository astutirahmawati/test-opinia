package utils

import (
	"opinia/configs"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlGorm(config *configs.ConfigSet) *gorm.DB {

	dsn := "root@tcp(localhost:3306)/opinia?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Warn(err)
	}
	return db
}

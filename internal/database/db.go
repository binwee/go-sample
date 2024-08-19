package database

import (
	"fmt"
	"log"

	"github.com/binwee/go-sample/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDb(cfg config.AppConfig) *gorm.DB {
	mysqlInfo := cfg.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", mysqlInfo.Username, mysqlInfo.Password,
		mysqlInfo.Ip, mysqlInfo.Port, mysqlInfo.DbName, mysqlInfo.Params)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	return db
}

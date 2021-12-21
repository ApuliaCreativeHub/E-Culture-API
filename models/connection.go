package models

import (
	"E-Culture-API/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var Db *gorm.DB

func InitializeDBConnection() {
	conf, err := utils.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	dsn := conf.Database.User + ":" + conf.Database.Password + "@tcp(" + conf.Database.Host + ":" +
		conf.Database.Port + ")/" + conf.Database.Schema + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})

	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"E-Culture-API/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB
var configPath = "conf/conf.yml"

func main() {
	err := utils.ValidateConfigPath(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	conf, err := utils.NewConfig(configPath)
	dsn := conf.Database.User + ":" + conf.Database.Password + "@tcp(" + conf.Database.Host + ":" +
		conf.Database.Port + ")/" + conf.Database.Schema + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
}

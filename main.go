package main

import (
	"E-Culture-API/models"
	"E-Culture-API/router"
	"E-Culture-API/utils"
	"log"
	"net/http"
	"time"
)

var configPath = "conf/conf.yml"

func main() {
	err := utils.SetConfigPath(configPath)
	if err != nil {
		log.Println(err)
		return
	}
	conf, err := utils.GetConfig()
	if err != nil {
		return
	}

	models.InitializeDBConnection()

	srv := &http.Server{
		Handler:      router.Router(),
		Addr:         conf.Server.Host + ":" + conf.Server.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

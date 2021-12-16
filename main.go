package main

import (
	"E-Culture-API/router"
	"E-Culture-API/utils"
	"log"
	"net/http"
	"time"
)

var configPath = "conf/conf.yml"

func main() {
	conf, err := utils.GetConfig(configPath)
	if err != nil {
		return
	}
	srv := &http.Server{
		Handler:      router.Router(),
		Addr:         conf.Server.Host + ":" + conf.Server.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

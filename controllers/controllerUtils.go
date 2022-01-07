package controllers

import (
	"E-Culture-API/models"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func checkAuthorization(r *http.Request) bool {
	authorizationHeader := r.Header.Get("Authorization")
	authorizationHeader = authorizationHeader[len("Bearer "):]
	t := new(models.Token)
	t.Token = authorizationHeader
	rows, err := t.ReadByToken()
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Println(err)
		}
		return false
	}

	return rows > 0
}

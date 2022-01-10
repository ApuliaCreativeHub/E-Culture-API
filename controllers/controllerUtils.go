package controllers

import (
	"E-Culture-API/models"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func checkAuthorization(r *http.Request) bool {
	authorizationHeader := getTokenFromHeader(r)
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

func getTokenFromHeader(r *http.Request) string {
	// TODO: Check if header is present and well formed
	authorizationHeader := r.Header.Get("Authorization")
	return authorizationHeader[len("Bearer "):]
}

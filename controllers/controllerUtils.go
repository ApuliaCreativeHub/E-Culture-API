package controllers

import (
	"E-Culture-API/models"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
)

func checkAuthorization(r *http.Request) bool {
	authorizationHeader, err := getTokenFromHeader(r)

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

func getTokenFromHeader(r *http.Request) (string, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" || !strings.Contains(authorizationHeader, "Bearer ") {
		return "", fmt.Errorf("no valid authorization header found")
	}
	return authorizationHeader[len("Bearer "):], nil
}

func getUserByToken(r *http.Request) (models.User, error) {
	strToken, err := getTokenFromHeader(r)
	tkn := models.Token{Token: strToken}
	_, err = tkn.ReadByToken()
	if err != nil {
		return models.User{}, err
	}
	user := models.User{ID: tkn.UserID}
	if !user.ReadAndIsACurator() {
		return models.User{}, err
	}
	return user, nil
}

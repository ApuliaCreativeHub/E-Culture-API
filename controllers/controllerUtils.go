package controllers

import (
	"E-Culture-API/controllers/utils"
	"E-Culture-API/models"
	"encoding/json"
	"errors"
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
	if user.ReadAndIsACurator() {
		user.IsACurator = true
	}
	return user, nil
}

func isUserAbleToAct(r *http.Request, structureUserId uint) error {
	user, err := getUserByToken(r)
	if err != nil {
		return err
	}

	if user.ID != structureUserId {
		return errors.New("UnauthorizedAction")
	}
	return nil
}

func setFileName(items interface{}) {
	switch items.(type) {
	case []models.Place:
		places := items.([]models.Place)
		for i := range places {
			places[i].NormalSizeImg = places[i].PhotoPath + "/" + places[i].FileName + "_n.png"
		}
		break
	case []models.Object:
		objects := items.([]models.Object)
		for i := range objects {
			objects[i].NormalSizeImg = objects[i].PhotoPath + "/" + objects[i].FileName + "_n.png"
		}
		break
	}
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) error {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		log.Println("Error while marshaling JSON...")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils.General5xx))
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonBody)
	if err != nil {
		log.Println("Error while sending Auth response...")
		return err
	}
	return nil
}

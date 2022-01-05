package controllers

import (
	"E-Culture-API/models"
	"E-Culture-API/utils"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// AddUser handles endpoint user/add
func AddUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = user.Create()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// AuthUser handles endpoint user/auth
func AuthUser(w http.ResponseWriter, r *http.Request) {
	uwt := models.UserWithToken{}
	err := json.NewDecoder(r.Body).Decode(&uwt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	uwt.User.Password, err = utils.CryptSHA1(uwt.User.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user := models.User{
		Email: uwt.User.Email,
	}
	err = user.ReadByEmail()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if user.Password == uwt.User.Password {
		token, err := utils.NewJSONWebToken()
		if err != nil {
			log.Println("Error while creating a new JWT...")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		t := models.Token{
			Token:     token,
			CreatedAt: time.Now(),
			UUID:      uwt.Token.UUID,
			UserID:    user.ID,
		}
		_ = t.Delete()

		err = t.Create()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		user.Password = ""
		uwt := models.UserWithToken{
			User:  user,
			Token: models.Token{Token: token},
		}

		jsonBody, err := json.Marshal(uwt)
		if err != nil {
			log.Println("Error while marshaling JSON...")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonBody)
		if err != nil {
			log.Println("Error while sending Auth response...")
			return
		}
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

// Logout handles endpoint user/logout
func Logout(w http.ResponseWriter, r *http.Request) {
	t := new(models.Token)
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Println("Error while unmarshaling JSON...")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = t.Delete()
}

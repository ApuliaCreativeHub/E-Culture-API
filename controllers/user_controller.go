package controllers

import (
	"E-Culture-API/models"
	"E-Culture-API/utils"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

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

func AuthUser(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		UUID     string `json:"uuid"`
	}
	cred := credentials{}
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cred.Password, err = utils.CryptSHA1(cred.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user := models.User{
		Email: cred.Email,
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

	if user.Password == cred.Password && user.Email == cred.Email {
		token, err := utils.NewJSONWebToken()
		if err != nil {
			log.Println("Error while creating a new JWT...")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		t := models.Token{
			Token:     token,
			CreatedAt: time.Now(),
			UUID:      cred.UUID,
			UserID:    user.ID,
		}
		_ = t.Delete()

		err = t.Create()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonMap := make(map[string]string)
		jsonMap["token"] = token
		jsonBody, err := json.Marshal(jsonMap)
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

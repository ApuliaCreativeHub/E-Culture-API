package controllers

import (
	"E-Culture-API/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := models.User{Email: vars["email"]}
	err := user.ReadByEmail()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func AuthUser(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	cred := credentials{}
	err := json.NewDecoder(r.Body).Decode(&cred)
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
		w.WriteHeader(http.StatusOK)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

}

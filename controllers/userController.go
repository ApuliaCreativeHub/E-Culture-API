package controllers

import (
	utils2 "E-Culture-API/controllers/utils"
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
		_, _ = w.Write([]byte(utils2.MalformedData))
		return
	}

	tempUser := models.User{Email: user.Email}
	err = tempUser.ReadByEmail()
	if tempUser.ID != 0 {
		w.WriteHeader(http.StatusConflict)
		_, _ = w.Write([]byte(utils2.EmailAlreadyExists))
		return
	}

	err = user.Create()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("50"))
		return
	}
}

// AuthUser handles endpoint user/auth
func AuthUser(w http.ResponseWriter, r *http.Request) {
	uwt := models.UserWithToken{}
	err := json.NewDecoder(r.Body).Decode(&uwt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils2.MalformedData))
		return
	}

	uwt.User.Password, err = utils.CryptSHA1(uwt.User.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils2.General5xx))
		return
	}
	user := models.User{
		Email: uwt.User.Email,
	}
	err = user.ReadByEmail()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils2.General5xx))
		return
	}

	if user.ID == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(utils2.EmailDoesNotExist))
		return
	}

	if user.Password == uwt.User.Password {
		token, err := utils.NewJSONWebToken()
		if err != nil {
			log.Println("Error while creating a new JWT...")
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils2.General5xx))
			return
		}
		t := models.Token{
			Token:     token,
			CreatedAt: time.Now(),
			UUID:      uwt.Token.UUID,
			UserID:    user.ID,
		}
		_ = t.DeleteByUUID()

		err = t.Create()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils2.General5xx))
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
			_, _ = w.Write([]byte(utils2.General5xx))
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
		_, _ = w.Write([]byte(utils2.IncorrectCredentials))
		return
	}
}

// Logout handles endpoint user/logout
func Logout(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r) {
		authorizationHeader, err := getTokenFromHeader(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		t := new(models.Token)
		t.Token = authorizationHeader
		_ = t.DeleteByToken()
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//UpdateUser handles endpoint user/update
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r) {
		var err error
		t := new(models.Token)
		t.Token, err = getTokenFromHeader(r)
		_, err = t.ReadByToken()
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		usrUpdated := new(models.User)
		err = json.NewDecoder(r.Body).Decode(usrUpdated)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils2.MalformedData))
			return
		}

		u := new(models.User)
		u = usrUpdated
		u.ID = t.UserID
		err = u.Update()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils2.UpdatingDataFailed))
			return
		}

		u.ID = 0
		u.Password = ""
		jsonBody, err := json.Marshal(u)
		if err != nil {
			log.Println("Error while marshaling JSON...")
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils2.General5xx))
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
	}
}

//ChangePassword handles endpoint user/changepassword
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils2.MalformedData))
		return
	}

	err = user.ReadByEmail()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils2.EmailDoesNotExist))
		return
	}
	tempPsw, err := utils.GenerateRandomString(8)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils2.RecoveringPasswordFailed))
		return
	}

	user.Password = tempPsw
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils2.RecoveringPasswordFailed))
		return
	}

	err = user.Update()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils2.RecoveringPasswordFailed))
		return
	}

	err = sendPasswordByEmail(tempPsw, user.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils2.RecoveringPasswordFailed))
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r) {
		user := models.User{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils2.MalformedData))
			return
		}

		err = user.ReadByEmail()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils2.General5xx))
			return
		}
		err = user.Delete()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils2.General5xx))
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func sendPasswordByEmail(password, email string) error {
	msg := "Subject: Recover password ECulture-Tool\n\n" +
		"This is your temporary password:\n" + password + "\nUse it to log in your account and then change it!"

	var emails []string
	emails = append(emails, email)
	err := utils2.SendMail(msg, emails)
	if err != nil {
		return err
	}

	return nil
}

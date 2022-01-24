package controllers

import (
	"E-Culture-API/controllers/utils"
	"E-Culture-API/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

//AddPlace handles endpoint place/add
func AddPlace(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r) {
		user, err := getUserByToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		place := new(models.Place)
		place.UserID = user.ID

		err = r.ParseMultipartForm(0)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.MalformedData))
			return
		}
		place.Name = r.PostFormValue("name")
		place.Address = r.PostFormValue("address")
		place.Description = r.PostFormValue("description")

		tempPlace := models.Place{Address: place.Address}
		err = tempPlace.ReadByAddress()
		if tempPlace.ID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(utils.PlaceAddressAlreadyExists))
			return
		}
		var ll utils.LatLong
		err = utils.RetrieveLatLong(place.Address, &ll)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(utils.AddressSearchFailed))
			return
		}
		place.Lat = ll.Lat
		place.Long = ll.Long

		photo, _, err := r.FormFile("img")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.ProcessingImagesFailed))
			return
		}

		path := "static/images/" + strconv.Itoa(int(place.ID))
		err = utils.MakeImgs(photo, path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}

		err = place.Create()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}

		place.PhotoPath = path
		err = place.Update()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetYourPlaces handles endpoint place/getYours
func GetYourPlaces(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r) {
		user, err := getUserByToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		place := new(models.Place)
		place.UserID = user.ID

		places, err := place.ReadByUserId()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.PlaceDoesNotExists))
			return
		}

		for i := range places {
			places[i].NormalSizeImg = places[i].PhotoPath + "/normal_size.png"
			places[i].Thumbnail = places[i].PhotoPath + "/thumbnail.png"
		}

		jsonBody, err := json.Marshal(places)
		if err != nil {
			log.Println("Error while marshaling JSON...")
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonBody)
		if err != nil {
			log.Println("Error while sending Auth response...")
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
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

func DeletePlace(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r) {
		place := models.Place{}
		err := json.NewDecoder(r.Body).Decode(&place)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.MalformedData))
			return
		}

		err = place.Read()
		if err != nil || place.ID == 0 {
			w.WriteHeader(http.StatusConflict)
			_, _ = w.Write([]byte(utils.PlaceDoesNotExists))
			return
		}

		user, err := getUserByToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if user.ID != place.UserID {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = place.Delete()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

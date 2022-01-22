package controllers

import (
	"E-Culture-API/controllers/utils"
	"E-Culture-API/models"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

//AddPlace handles endpoint place/add
func AddPlace(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r) {
		place := new(models.Place)
		//TODO: change multipart/form-data to application/json and use base64 encoding for images
		placeJSON, _, err := r.FormFile("place")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(utils.MalformedData))
			return
		}
		err = json.NewDecoder(placeJSON).Decode(&place)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(utils.MalformedData))
			return
		}

		strToken, err := getTokenFromHeader(r)
		tkn := models.Token{Token: strToken}
		_, err = tkn.ReadByToken()
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		user := models.User{ID: tkn.UserID}
		if !user.ReadAndIsACurator() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

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
			_, _ = w.Write([]byte(utils.AddressSearchingFailed))
			return
		}
		place.Lat = ll.Lat
		place.Long = ll.Long

		photo, _, err := r.FormFile("img")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(utils.MalformedData))
			return
		}
		all, err := io.ReadAll(photo)
		if err != nil {
			return
		}
		//This is just a test
		err = ioutil.WriteFile("static/images/test.jpg", all, 0655)
		if err != nil {
			return
		}
		err = place.Create()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.AddingPlaceFailed))
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

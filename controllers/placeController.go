package controllers

import (
	"E-Culture-API/controllers/utils"
	"E-Culture-API/models"
	"net/http"
	"strconv"
)

//AddPlace handles endpoint place/add
func AddPlace(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r) {
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
			_, _ = w.Write([]byte(utils.AddressSearchingFailed))
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
			_, _ = w.Write([]byte(utils.AddingPlaceFailed))
			return
		}
		place.PhotoPath = path
		err = place.Update()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.AddingPlaceFailed))
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

package controllers

import (
	"E-Culture-API/controllers/utils"
	"E-Culture-API/models"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
)

var imgPlacesPath = "static/images/place/"

//AddPlace handles endpoint place/add
func AddPlace(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r) {
		place, photo, err := retrieveMultipartPlace(w, r)
		if err != nil {
			return
		}
		err = place.Create()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}

		path := imgPlacesPath + strconv.Itoa(int(place.ID))
		fileName, err := utils.MakeImgs(photo, path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}

		place.PhotoPath = path
		place.FileName = fileName
		err = place.Update()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}

		_ = sendJSONResponse(w, place)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func retrieveMultipartPlace(w http.ResponseWriter, r *http.Request) (*models.Place, multipart.File, error) {
	user, err := getUserByToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return nil, nil, err
	}
	place := new(models.Place)
	place.UserID = user.ID

	err = r.ParseMultipartForm(0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils.MalformedData))
		return nil, nil, err
	}
	place.Name = r.PostFormValue("name")
	place.Address = r.PostFormValue("address")
	place.Description = r.PostFormValue("description")
	uintID, err := strconv.Atoi(r.PostFormValue("id"))
	if err == nil {
		place.ID = uint(uintID)
	}

	tempPlace := models.Place{Address: place.Address}
	err = tempPlace.ReadByAddress()
	if tempPlace.ID != place.ID && tempPlace.ID != 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(utils.PlaceAddressAlreadyExists))
		return nil, nil, fmt.Errorf("address already exists")
	}
	var ll utils.LatLong
	err = utils.RetrieveLatLong(place.Address, &ll)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(utils.AddressSearchFailed))
		return nil, nil, err
	}
	place.Lat = ll.Lat
	place.Long = ll.Long

	photo, _, _ := r.FormFile("img")

	return place, photo, nil
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

		setFileName(places)

		err = sendJSONResponse(w, places)
		if err != nil {
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//DeletePlace handles endpoint place/delete
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

		err = isUserAbleToAct(r, place.UserID)
		if err != nil {
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

// UpdatePlace handles endpoint place/update
func UpdatePlace(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r) {
		place, photo, err := retrieveMultipartPlace(w, r)
		if err != nil {
			return
		}

		err = isUserAbleToAct(r, place.UserID)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if photo != nil {
			path := imgPlacesPath + strconv.Itoa(int(place.ID))
			fileName, err := utils.MakeImgs(photo, path)
			place.FileName = fileName
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(utils.General5xx))
				return
			}
			place.PhotoPath = path
		}

		err = place.Update()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}

		_ = sendJSONResponse(w, place)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetPlaces handles endpoint place/get
func GetPlaces(w http.ResponseWriter, _ *http.Request) {
	place := new(models.Place)
	places, err := place.ReadAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	setFileName(places)

	_ = sendJSONResponse(w, places)
}

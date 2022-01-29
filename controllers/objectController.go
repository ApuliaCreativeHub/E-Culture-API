package controllers

import (
	"E-Culture-API/controllers/utils"
	"E-Culture-API/models"
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
)

//AddObject handles endpoint object/add
func AddObject(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r) {
		object, photo, err := retrieveMultipartObject(w, r)
		if err != nil {
			return
		}
		err = object.Create()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}

		path := "static/images/object/" + strconv.Itoa(int(object.ID))
		fileName, err := utils.MakeImgs(photo, path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}

		object.PhotoPath = path
		object.FileName = fileName
		err = object.Update()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}

		if sendObjectJson(w, object) != nil {
			return
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetZoneObjects handles endpoint object/getZoneObjects
func GetZoneObjects(w http.ResponseWriter, r *http.Request) {
	zoneId, err := strconv.Atoi(r.FormValue("zoneId"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils.General5xx))
		return
	}

	object := models.Object{ZoneID: uint(zoneId)}
	objects, err := object.ReadByZoneId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils.General5xx))
		return
	}

	setFileName(objects)

	jsonBody, err := json.Marshal(objects)
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
}

// UpdateObject handles endpoint object/update
func UpdateObject(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r) {
		object, photo, err := retrieveMultipartObject(w, r)
		if err != nil {
			return
		}

		tempObj := *object
		err = tempObj.ReadAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}
		err = isUserAbleToAct(r, tempObj.Zone.Place.UserID)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if photo != nil {
			path := "static/images/object/" + strconv.Itoa(int(object.ID))
			fileName, err := utils.MakeImgs(photo, path)
			object.FileName = fileName
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(utils.General5xx))
				return
			}
			object.PhotoPath = path
		}

		err = object.Update()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}

		if sendObjectJson(w, object) != nil {
			return
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func DeleteObject(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r) {
		object := models.Object{}
		err := json.NewDecoder(r.Body).Decode(&object)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.MalformedData))
			return
		}

		err = object.ReadAll()
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			_, _ = w.Write([]byte(utils.ObjectDoesNotExists))
			return
		}

		err = isUserAbleToAct(r, object.Zone.Place.UserID)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = object.Delete()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func sendObjectJson(w http.ResponseWriter, object *models.Object) error {
	jsonBody, err := json.Marshal(object)
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

func retrieveMultipartObject(w http.ResponseWriter, r *http.Request) (*models.Object, multipart.File, error) {
	object := new(models.Object)

	err := r.ParseMultipartForm(0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils.MalformedData))
		return nil, nil, err
	}
	object.Name = r.PostFormValue("name")
	ZoneId, err := strconv.Atoi(r.PostFormValue("zoneId"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils.MalformedData))
		return nil, nil, err
	}
	object.ZoneID = uint(ZoneId)
	object.Description = r.PostFormValue("description")
	uintID, err := strconv.Atoi(r.PostFormValue("id"))
	if err == nil {
		object.ID = uint(uintID)
	}

	photo, _, _ := r.FormFile("img")

	return object, photo, nil
}

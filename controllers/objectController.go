package controllers

import (
	"E-Culture-API/controllers/utils"
	"E-Culture-API/models"
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

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
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

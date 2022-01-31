package controllers

import (
	"E-Culture-API/controllers/utils"
	"E-Culture-API/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type PathWithObjects struct {
	Path    models.Path     `json:"path"`
	Objects []models.Object `json:"objects"`
}

func AddPath(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r) {
		pwo := PathWithObjects{}
		err := json.NewDecoder(r.Body).Decode(&pwo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.MalformedData))
			return
		}

		user, err := getUserByToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		pwo.Path.UserID = user.ID
		err = pwo.Path.Create()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}

		for i, o := range pwo.Objects {
			err = pwo.Path.AddObjectToPath(o.ID, uint(i))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(utils.General5xx))
				return
			}
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func GetPlacePaths(w http.ResponseWriter, r *http.Request) {
	placeId, err := strconv.Atoi(r.FormValue("placeId"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils.General5xx))
		return
	}

	path := models.Path{}
	paths, err := path.ReadPathsByPlaceId(uint(placeId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils.General5xx))
		return
	}

	_ = sendJSONResponse(w, paths)
}

func GetUserPaths(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.FormValue("userId"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils.General5xx))
		return
	}

	path := models.Path{}
	paths, err := path.ReadByUserId(uint(userId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils.General5xx))
		return
	}

	_ = sendJSONResponse(w, paths)
}

func UpdatePath(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r) {
		pwo := PathWithObjects{}
		err := json.NewDecoder(r.Body).Decode(&pwo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.MalformedData))
			return
		}

		tempPath := pwo.Path
		err = tempPath.ReadByPathId()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.ZoneNameAlreadyExists))
			return
		}

		err = isUserAbleToAct(r, tempPath.UserID)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = pwo.Path.Update(pwo.Objects)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func DeletePath(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r) {
		path := models.Path{}
		err := json.NewDecoder(r.Body).Decode(&path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.MalformedData))
			return
		}

		tempPath := path
		err = tempPath.ReadByPathId()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.ZoneNameAlreadyExists))
			return
		}

		err = isUserAbleToAct(r, tempPath.UserID)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = path.Delete()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

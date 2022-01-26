package controllers

import (
	"E-Culture-API/controllers/utils"
	"E-Culture-API/models"
	"encoding/json"
	"net/http"
)

// AddZone handles endpoint zone/add
func AddZone(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r) {
		zone := models.Zone{}
		err := json.NewDecoder(r.Body).Decode(&zone)
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

		place := models.Place{ID: zone.PlaceID}
		err = place.Read()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}
		if place.UserID != user.ID {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = zone.ReadByName()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.ZoneNameAlreadyExists))
			return
		}

		err = zone.Create()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

package controllers

import (
	"E-Culture-API/controllers/utils"
	"E-Culture-API/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

// GetPlaceZones handles endpoint zone/getPlaceZones
func GetPlaceZones(w http.ResponseWriter, r *http.Request) {
	placeId, err := strconv.Atoi(r.FormValue("placeId"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils.General5xx))
		return
	}

	zone := models.Zone{PlaceID: uint(placeId)}
	zones, err := zone.ReadByPlaceId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(utils.General5xx))
		return
	}

	jsonBody, err := json.Marshal(zones)
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

func DeleteZone(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r) {
		zone := models.Zone{}
		err := json.NewDecoder(r.Body).Decode(&zone)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.MalformedData))
			return
		}

		err = zone.ReadAndPreloadPlace()
		if err != nil || zone.ID == 0 {
			w.WriteHeader(http.StatusConflict)
			_, _ = w.Write([]byte(utils.ZoneDoesNotExists))
			return
		}

		err = isUserAbleToAct(r, zone.Place.UserID)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = zone.Delete()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func UpdateZone(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r) {
		zone := models.Zone{}
		err := json.NewDecoder(r.Body).Decode(&zone)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.MalformedData))
			return
		}

		tempZone := zone
		err = tempZone.ReadAndPreloadPlace()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}
		err = isUserAbleToAct(r, tempZone.Place.UserID)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// TODO: Check that a zone with same name for the same place doesn't exist first!
		err = zone.Update()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(utils.General5xx))
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

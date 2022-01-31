package router

import (
	"E-Culture-API/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	// static router: serve files under /static/<filename>
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// User routers
	r.HandleFunc("/user/auth", controllers.AuthUser).Methods("POST")
	r.HandleFunc("/user/add", controllers.AddUser).Methods("POST")
	r.HandleFunc("/user/logout", controllers.Logout).Methods("GET")
	r.HandleFunc("/user/update", controllers.UpdateUser).Methods("POST")
	r.HandleFunc("/user/delete", controllers.DeleteUser).Methods("POST")
	r.HandleFunc("/user/changepassword", controllers.ChangePassword).Methods("POST")

	// Place routers
	r.HandleFunc("/place/add", controllers.AddPlace).Methods("POST")
	r.HandleFunc("/place/getYours", controllers.GetYourPlaces).Methods("GET")
	r.HandleFunc("/place/getAll", controllers.GetPlaces).Methods("GET")
	r.HandleFunc("/place/update", controllers.UpdatePlace).Methods("POST")
	r.HandleFunc("/place/delete", controllers.DeletePlace).Methods("POST")

	//Zone routers
	r.HandleFunc("/zone/add", controllers.AddZone).Methods("POST")
	r.Path("/zone/getPlaceZones").Queries("placeId", "{placeId}").HandlerFunc(controllers.GetPlaceZones).Methods("GET")
	r.HandleFunc("/zone/delete", controllers.DeleteZone).Methods("POST")
	r.HandleFunc("/zone/update", controllers.UpdateZone).Methods("POST")

	//Object routers
	r.HandleFunc("/object/add", controllers.AddObject).Methods("POST")
	r.Path("/object/getZoneObjects").Queries("zoneId", "{zoneId}").HandlerFunc(controllers.GetZoneObjects).Methods("GET")
	r.Path("/object/getById").Queries("id", "{id}").HandlerFunc(controllers.GetObjectById).Methods("GET")
	r.HandleFunc("/object/update", controllers.UpdateObject).Methods("POST")
	r.HandleFunc("/object/delete", controllers.DeleteObject).Methods("POST")

	// Path routers
	r.HandleFunc("/path/add", controllers.AddPath).Methods("POST")
	r.Path("/path/getPlacePaths").Queries("placeId", "{placeId}").HandlerFunc(controllers.GetPlacePaths).Methods("GET")
	r.Path("/path/getUserPaths").Queries("userId", "{userId}").HandlerFunc(controllers.GetUserPaths).Methods("GET")
	r.HandleFunc("/path/update", controllers.UpdatePath).Methods("POST")
	r.HandleFunc("/path/delete", controllers.DeletePath).Methods("POST")

	return r
}

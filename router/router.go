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

	return r
}

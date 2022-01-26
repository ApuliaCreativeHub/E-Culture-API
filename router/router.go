package router

import (
	"E-Culture-API/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	//static route: serve files under /static/<filename>
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//User router
	r.HandleFunc("/user/auth", controllers.AuthUser).Methods("POST")
	r.HandleFunc("/user/add", controllers.AddUser).Methods("POST")
	r.HandleFunc("/user/logout", controllers.Logout).Methods("GET")
	r.HandleFunc("/user/update", controllers.UpdateUser).Methods("POST")
	r.HandleFunc("/user/delete", controllers.DeleteUser).Methods("POST")
	r.HandleFunc("/user/changepassword", controllers.ChangePassword).Methods("POST")
	//Place router
	r.HandleFunc("/place/add", controllers.AddPlace).Methods("POST")
	r.HandleFunc("/place/getYours", controllers.GetYourPlaces).Methods("GET")
	r.HandleFunc("/place/getAll", controllers.GetPlaces).Methods("GET")
	r.HandleFunc("/place/update", controllers.UpdatePlace).Methods("POST")
	r.HandleFunc("/place/delete", controllers.DeletePlace).Methods("POST")
	return r
}

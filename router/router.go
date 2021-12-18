package router

import (
	"E-Culture-API/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	//TODO: add other routers
	r.HandleFunc("/user/auth", controllers.AuthUser)
	return r
}

package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	//Router example
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("HelloWorld!")) })
	//TODO: add other routers
	return r
}

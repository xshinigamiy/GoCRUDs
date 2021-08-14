package routes

import (
	"GoCRUDs/internal/controller"
	"github.com/gorilla/mux"
)

func Route(r *mux.Router) {
	r.HandleFunc("/create-user", controller.CreateUser).Methods("POST")
}

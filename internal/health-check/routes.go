package health_check

import (
	"github.com/gorilla/mux"
)

type Health struct{}

func HealthCheckRoutes(r *mux.Router) {
	h := Health{}
	r.HandleFunc("/health-check", h.CheckHealth).Methods("GET")
}

package user

import (
	"github.com/gorilla/mux"
)

// SetRoutes add routes from user
func SetRoutes(r *mux.Router) {
	r.HandleFunc("", Add).Methods("POST")
	r.HandleFunc("", GetAll).Methods("GET")
	r.HandleFunc("/fb/{id}", FindByFacebookID).Methods("GET")
	r.HandleFunc("/{id}", FindByID).Methods("GET")
	r.HandleFunc("/{id}", DeleteByID).Methods("DELETE")
	r.HandleFunc("/{id}", UpdateByID).Methods("PUT")
}

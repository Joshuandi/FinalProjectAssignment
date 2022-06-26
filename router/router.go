package router

import (
	photo_handler "FinalProjectAssignment/handler"
	user_handler "FinalProjectAssignment/handler"
	"FinalProjectAssignment/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func UserRouter(r *mux.Router, u *user_handler.UserHandler) {
	r.HandleFunc("/users/register", u.UserRegister).Methods("POST")
	r.HandleFunc("/users/login", u.UserLogin).Methods("POST")
	r.Handle("/users", middleware.Authorization(http.HandlerFunc(u.UserUpdate))).Methods("PUT")
	r.Handle("/users/account", middleware.Authorization(http.HandlerFunc(u.UserGetId))).Methods("GET")
	r.Handle("/users", middleware.Authorization(http.HandlerFunc(u.UserDelete))).Methods("DELETE")
}

func PhotoRouter(r *mux.Router, p *photo_handler.PhotoHandler) {
	r.Handle("/photos", middleware.Authorization(http.HandlerFunc(p.PhotoRegister))).Methods("POST")
	r.Handle("/photos", middleware.Authorization(http.HandlerFunc(p.PhotoGet))).Methods("GET")
	r.Handle("/photos/{id}", middleware.Authorization(http.HandlerFunc(p.PhotoUpdate))).Methods("PUT")
	r.Handle("/photos/{id}", middleware.Authorization(http.HandlerFunc(p.PhotoDelete))).Methods("DELETE")
}

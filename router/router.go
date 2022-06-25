package router

import (
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

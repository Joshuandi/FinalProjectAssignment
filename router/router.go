package router

import (
	user_handler "FinalProjectAssignment/handler"

	"github.com/gorilla/mux"
)

func UserRouter(r *mux.Router, u *user_handler.UserHandler) {
	r.HandleFunc("/users/register", u.UserRegister).Methods("POST")
	r.HandleFunc("/users/login", u.UserLogin).Methods("POST")
	//r.HandleFunc("/users", middleware.AuthIsAuthorized(http.HandlerFunc(u.UserUpdate))).Methods("PUT")
	//r.HandleFunc("/users/profile", middleware.AuthIsAuthorized(http.HandlerFunc(u.UserGetId))).Methods("GET")
	//r.HandleFunc("/users", middleware.AuthIsAuthorized(http.HandlerFunc(u.UserDelete))).Methods("DELETE")
}

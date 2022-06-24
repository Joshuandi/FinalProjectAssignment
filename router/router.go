package router

import (
	user_handler "FinalProjectAssignment/handler"

	"github.com/gorilla/mux"
)

func UserRouter(r *mux.Router, u *user_handler.UserHandler) {
	r.HandleFunc("/users/register", u.UserRegister).Methods("POST")
	r.HandleFunc("/users/login", u.UserLogin).Methods("POST")
}

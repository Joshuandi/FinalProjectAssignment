package router

import (
	"FinalProjectAssignment/handler"
	"FinalProjectAssignment/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func UserRouter(r *mux.Router, u *handler.UserHandler) {
	r.HandleFunc("/users/register", u.UserRegister).Methods("POST")
	r.HandleFunc("/users/login", u.UserLogin).Methods("POST")
	r.Handle("/users", middleware.Authorization(http.HandlerFunc(u.UserUpdate))).Methods("PUT")
	r.Handle("/users/account", middleware.Authorization(http.HandlerFunc(u.UserGetId))).Methods("GET")
	r.Handle("/users", middleware.Authorization(http.HandlerFunc(u.UserDelete))).Methods("DELETE")
}

func PhotoRouter(r *mux.Router, p *handler.PhotoHandler) {
	r.Handle("/photos", middleware.Authorization(http.HandlerFunc(p.PhotoRegister))).Methods("POST")
	r.Handle("/photos", middleware.Authorization(http.HandlerFunc(p.PhotoGet))).Methods("GET")
	r.Handle("/photos/{id}", middleware.Authorization(http.HandlerFunc(p.PhotoUpdate))).Methods("PUT")
	r.Handle("/photos/{id}", middleware.Authorization(http.HandlerFunc(p.PhotoDelete))).Methods("DELETE")
}

func CommentRouter(r *mux.Router, c *handler.CommentHandler) {
	r.Handle("/comments", middleware.Authorization(http.HandlerFunc(c.CommentRegister))).Methods("POST")
	r.Handle("/comments", middleware.Authorization(http.HandlerFunc(c.CommentGet))).Methods("GET")
	r.Handle("/comments/{id}", middleware.Authorization(http.HandlerFunc(c.CommentUpdate))).Methods("PUT")
	r.Handle("/comments/{id}", middleware.Authorization(http.HandlerFunc(c.CommentDelete))).Methods("DELETE")
}

func Social_MediaRouter(r *mux.Router, sm *handler.SocialMediaHandler) {
	r.Handle("/socialmedias", middleware.Authorization(http.HandlerFunc(sm.SocialMediaRegister))).Methods("POST")
	r.Handle("/socialmedias", middleware.Authorization(http.HandlerFunc(sm.SocialMediaGet))).Methods("GET")
	r.Handle("/socialmedias/{id}", middleware.Authorization(http.HandlerFunc(sm.SocialMediaUpdate))).Methods("PUT")
	r.Handle("/socialmedias/{id}", middleware.Authorization(http.HandlerFunc(sm.SocialMediaDelete))).Methods("DELETE")
}

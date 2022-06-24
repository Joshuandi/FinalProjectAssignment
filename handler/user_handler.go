package handler

import (
	"FinalProjectAssignment/auth"
	"FinalProjectAssignment/model"
	"FinalProjectAssignment/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type UserHandlerInterface interface{}
type UserHandler struct {
	userSrvc service.UserServiceInterface
}

func NewUserhandler(userSrvc service.UserServiceInterface) *UserHandler {
	return &UserHandler{userSrvc: userSrvc}
}

func (u *UserHandler) UserRegister(w http.ResponseWriter, r *http.Request) {
	var users *model.User
	//read dari body json
	jsonDec := json.NewDecoder(r.Body)
	err := jsonDec.Decode(&users)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//throw to service
	newRegister, err := u.userSrvc.UserRegister(r.Context(), &model.User{
		Username: users.Username,
		Password: users.Password,
		Email:    users.Email,
		Age:      users.Age,
	})
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//thorw for respone result
	Register_respone := model.UserRegisterRespone{
		R_user_id:  newRegister.User_id,
		R_email:    newRegister.Email,
		R_username: newRegister.Username,
		R_age:      newRegister.Age,
	}
	jsonData, _ := json.Marshal(Register_respone)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)

	fmt.Println("ini respone", Register_respone)
	return
}
func (u *UserHandler) UserLogin(w http.ResponseWriter, r *http.Request) {
	var login *model.UserPostLogin
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("ini login handler: ", login)
	newLogin, err := u.userSrvc.UserLogin(r.Context(), login)
	fmt.Println("ini login handler: ", newLogin)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("ini newlogin handlder: ", newLogin)
	id := strconv.Itoa(newLogin.User_id)
	validToken, err := auth.GenerateJWT(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Token not Valid", http.StatusBadRequest)
		return
	}
	var token model.UserToken
	token.TokenString = validToken
	fmt.Println("ini token :", token)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)

}

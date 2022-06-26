package handler

import (
	"FinalProjectAssignment/auth"
	"FinalProjectAssignment/middleware"
	"FinalProjectAssignment/model"
	"FinalProjectAssignment/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

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
	w.WriteHeader(201)
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
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(token)
}

func (u *UserHandler) UserUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.RunUser(ctx)
	fmt.Println("userid update :", user.User_id)

	var userUp *model.UserUpdateInput
	json.NewDecoder(r.Body).Decode(&userUp)

	id := strconv.Itoa(user.User_id)
	fmt.Println("userid update :", id)
	userUpdate, err := u.userSrvc.UserUpdate(r.Context(), id, &model.UserUpdateInput{
		U_email:      userUp.U_email,
		U_username:   userUp.U_username,
		U_Updated_at: userUp.U_Updated_at,
	})
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	Update_respone := model.UserUpdateRespone{
		U_user_id:    user.User_id,
		U_email:      userUpdate.Email,
		U_username:   userUpdate.Username,
		U_age:        userUpdate.Age,
		U_Updated_at: user.Updated_at,
	}
	jsonData, _ := json.Marshal(Update_respone)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonData)

	fmt.Println("ini respone", Update_respone)
	return
}

func (u *UserHandler) UserDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.RunUser(ctx)
	fmt.Println("userid delete :", user.User_id)
	id := strconv.Itoa(user.User_id)
	fmt.Println("userid delete :", id)
	_, err := u.userSrvc.UserDelete(r.Context(), id, &model.User{
		User_id: user.User_id,
	})
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	msg := model.DeleteData{
		Message: "Your account has been successfully deleted",
	}
	w.Write([]byte(msg.Message))
	return
}

func (u *UserHandler) UserGetId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//fmt.Println("user handler", ctx)
	user := middleware.RunUser(ctx)
	id := strconv.Itoa(user.User_id)
	user_, err := u.userSrvc.UserGetId(ctx, id)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userUpdate := model.UserUpdateRespone{
		U_user_id:    user_.User_id,
		U_username:   user_.Username,
		U_email:      user_.Email,
		U_age:        user_.Age,
		U_Updated_at: user.Updated_at,
	}
	res, _ := json.Marshal(userUpdate)
	w.Header().Add("Content-Type", "application/json")
	w.Write(res)
	return
}

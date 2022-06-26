package model

import "time"

type User struct {
	User_id    int       `json:"user_id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Age        int       `json:"age"`
	Created_at time.Time `json:"create_at"`
	Updated_at time.Time `json:"updated_at"`
}

type UserPostLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserToken struct {
	TokenString string `json:"token"`
}

type UserRes struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResT struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserRegisterRespone struct {
	R_user_id  int    `json:"user_id"`
	R_username string `json:"username"`
	R_email    string `json:"email"`
	R_age      int    `json:"age"`
}

type UserUpdateInput struct {
	U_email      string    `json:"email"`
	U_username   string    `json:"username"`
	U_Updated_at time.Time `json:"updated_at"`
}

type UserUpdateRespone struct {
	U_user_id    int       `json:"user_id"`
	U_email      string    `json:"email"`
	U_username   string    `json:"username"`
	U_age        int       `json:"age"`
	U_Updated_at time.Time `json:"updated_at"`
}

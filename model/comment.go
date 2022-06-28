package model

import "time"
//
type Comment struct {
	Comment_id int       `json:"id"`
	Message    string    `json:"message"`
	User_id    int       `json:"user_id"`
	Photo_id   int       `json:"photo_id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type CommentRegisterRespone struct {
	R_Comment_id int       `json:"id"`
	R_Message    string    `json:"message"`
	R_User_id    int       `json:"user_id"`
	R_Photo_id   int       `json:"photo_id"`
	R_Created_at time.Time `json:"created_at"`
}

type CommentUpdateRespone struct {
	U_Comment_id int       `json:"id"`
	U_Message    string    `json:"message"`
	U_User_id    int       `json:"user_id"`
	U_Photo_id   int       `json:"photo_id"`
	U_Updated_at time.Time `json:"updated_at"`
}

type CommentGet struct {
	Comment_id int          `json:"id"`
	Message    string       `json:"message"`
	Photo_id   int          `json:"photo_id"`
	User_id    int          `json:"user_id"`
	Updated_at time.Time    `json:"updated_at"`
	Created_at time.Time    `json:"created_at"`
	User       UserComment  `json:"User"`
	Photo      PhotoComment `json:"Photo"`
}

type CommentShow struct {
	Comment_id int       `json:"id"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	Photo_url  string    `json:"photo_url"`
	User_id    int       `json:"user_id"`
	Updated_at time.Time `json:"updated_at"`
}

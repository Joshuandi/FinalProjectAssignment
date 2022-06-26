package model

import "time"

type Comment struct {
	Comment_id int       `json:"comment_id"`
	User_id    int       `json:"User_id"`
	Photo_id   int       `json:"Photo_id"`
	Message    string    `json:"message"`
	Created_at time.Time `json:"create_at"`
	Updated_at time.Time `json:"updated_at"`
	User       User      `json:"User"`
	Photo      Photo     `json:"Photo"`
}

type CommentRegisterRespone struct {
	R_Comment_id int       `json:"comment_id"`
	R_User_id    int       `json:"User_id"`
	R_Photo_id   int       `json:"Photo_id"`
	R_Message    string    `json:"message"`
	R_Created_at time.Time `json:"create_at"`
}

type CommentUpdateRespone struct {
	U_Comment_id int       `json:"comment_id"`
	U_User_id    int       `json:"User_id"`
	U_Photo_id   int       `json:"Photo_id"`
	U_Message    string    `json:"message"`
	U_Updated_at time.Time `json:"updated_at"`
}

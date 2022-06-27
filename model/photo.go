package model

import "time"

type Photo struct {
	Photo_id   int       `json:"id"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	Photo_url  string    `json:"photo_url"`
	User_id    int       `json:"User_id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type PhotoGet struct {
	Photo_id   int       `json:"id"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	Photo_url  string    `json:"photo_url"`
	User_id    int       `json:"user_id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	User       UserResT  `json:"User"`
}

type PhotoUpdateInput struct {
	U_photo_id   int       `json:"id"`
	U_title      string    `json:"title"`
	U_caption    string    `json:"caption"`
	U_photo_url  string    `json:"photo_url"`
	U_user_id    int       `json:"user_id"`
	U_updated_at time.Time `json:"updated_at"`
}

type PhotoRegisterRespone struct {
	R_photo_id   int       `json:"id"`
	R_title      string    `json:"title"`
	R_caption    string    `json:"caption"`
	R_photo_url  string    `json:"photo_url"`
	R_user_id    int       `json:"user_id"`
	R_created_at time.Time `json:"created_at"`
}

type PhotoUpdateRespone struct {
	U_photo_id   int       `json:"id"`
	U_title      string    `json:"title"`
	U_caption    string    `json:"caption"`
	U_photo_url  string    `json:"photo_url"`
	U_user_id    int       `json:"user_id"`
	U_updated_at time.Time `json:"updated_at"`
}

type PhotoComment struct {
	Photo_id  int    `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url"`
	User_id   int    `json:"user_id"`
}

type PhotoCommentShow struct {
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url"`
	User_id   int    `json:"user_id"`
}

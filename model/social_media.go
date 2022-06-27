package model

import "time"

type SocialMedia struct {
	Sm_Id           int       `json:"id"`
	Name            string    `json:"name"`
	SocialMedia_url string    `json:"social_media_url"`
	User_id         int       `json:"user_id"`
	Created_at      time.Time `json:"created_at"`
	Updated_at      time.Time `json:"updated_at"`
}

type SocialMediaRegisterRespone struct {
	R_Sm_Id           int       `json:"id"`
	R_Name            string    `json:"name"`
	R_SocialMedia_url string    `json:"social_media_url"`
	R_User_id         int       `json:"user_id"`
	R_Created_at      time.Time `json:"created_at"`
}

type SocialMediaUpdateRespone struct {
	U_Sm_Id           int       `json:"id"`
	U_Name            string    `json:"name"`
	U_SocialMedia_url string    `json:"social_media_url"`
	U_Updated_at      time.Time `json:"updated_at"`
	U_User_id         int       `json:"user_id"`
	U_User            User      `json:"User"`
}

type SocialMediaShow struct {
	Social_medias SocialMediaGet `json:"social_medias"`
}

type SocialMediaGet struct {
	Sm_Id           int             `json:"id"`
	Name            string          `json:"name"`
	SocialMedia_url string          `json:"social_media_url"`
	Created_at      time.Time       `json:"created_at"`
	Updated_at      time.Time       `json:"updated_at"`
	User_id         int             `json:"id"`
	User            UserSocialMedia `json:"User"`
}

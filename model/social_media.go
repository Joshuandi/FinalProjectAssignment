package model

import "time"

type SocialMedia struct {
	Sm_Id           int       `json:"socialMedia_id"`
	Name            string    `json:"name"`
	SocialMedia_url string    `json:"socialMedia_url"`
	Created_at      time.Time `json:"create_at"`
	Updated_at      time.Time `json:"updated_at"`
	User_id         int       `json:"User_id"`
	User            User      `json:"User"`
}

type SocialMediaRegisterRespone struct {
	R_Sm_Id           int       `json:"socialMedia_id"`
	R_Name            string    `json:"name"`
	R_SocialMedia_url string    `json:"socialMedia_url"`
	R_Created_at      time.Time `json:"create_at"`
	R_User_id         int       `json:"User_id"`
	R_User            User      `json:"User"`
}

type SocialMediaUpdateRespone struct {
	U_Sm_Id           int       `json:"socialMedia_id"`
	U_Name            string    `json:"name"`
	U_SocialMedia_url string    `json:"socialMedia_url"`
	U_Updated_at      time.Time `json:"updated_at"`
	U_User_id         int       `json:"User_id"`
	U_User            User      `json:"User"`
}

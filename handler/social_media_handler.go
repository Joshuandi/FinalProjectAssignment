package handler

import (
	"FinalProjectAssignment/middleware"
	"FinalProjectAssignment/model"
	"FinalProjectAssignment/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type SocialMediaHandler struct {
	SocialMediaSrvc service.SocialMediaServiceInterface
}

func NewSocialMediahandler(SocialMediaSrvc service.SocialMediaServiceInterface) *SocialMediaHandler {
	return &SocialMediaHandler{SocialMediaSrvc: SocialMediaSrvc}
}

func (s *SocialMediaHandler) SocialMediaRegister(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := middleware.RunUser(ctx)
	fmt.Println("ini SocialMedia handler : ", userId.User_id)
	var sm *model.SocialMedia
	//read dari body json
	jsonDec := json.NewDecoder(r.Body)
	err := jsonDec.Decode(&sm)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("ini SocialMedia handler : ", sm)
	//id := strconv.Itoa(userId.User_id)
	//throw to service
	newRegister, err := s.SocialMediaSrvc.SocialMediaRegister(r.Context(), &model.SocialMedia{
		Name:            sm.Name,
		SocialMedia_url: sm.SocialMedia_url,
		User_id:         userId.User_id,
		Created_at:      sm.Created_at,
	})
	fmt.Println("ini SocialMedia handler : ", newRegister)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//thorw for respone result
	Register_respone := model.SocialMediaRegisterRespone{
		R_Sm_Id:           newRegister.Sm_Id,
		R_Name:            newRegister.Name,
		R_SocialMedia_url: newRegister.SocialMedia_url,
		R_User_id:         newRegister.User_id,
		R_Created_at:      newRegister.Created_at,
	}
	fmt.Println("ini SocialMedia handler : ", Register_respone)
	jsonData, _ := json.Marshal(Register_respone)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(jsonData)

	fmt.Println("ini respone", Register_respone)
	return
}

func (s *SocialMediaHandler) SocialMediaGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ids := params["id"]
	fmt.Println(ids)
//
	ctx := r.Context()
	userId := middleware.RunUser(ctx)

	fmt.Println("ini SocialMedia handler : ", userId.User_id)
	SocialMedia, err := s.SocialMediaSrvc.SocialMediaGet(r.Context())
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(&SocialMedia)
	w.Header().Add("Content-Type", "application/json")
	w.Write(res)
	return
}

func (s *SocialMediaHandler) SocialMediaUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var sm *model.SocialMedia
	//read dari body json
	jsonDec := json.NewDecoder(r.Body)
	err := jsonDec.Decode(&sm)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("ini photo handler : ", sm)
	newUpdate, err := s.SocialMediaSrvc.SocialMediaUpdate(r.Context(), id, &model.SocialMedia{
		Name:            sm.Name,
		SocialMedia_url: sm.SocialMedia_url,
	})
	fmt.Println("ini photo handler : ", newUpdate)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	Update_respone := model.SocialMedia{
		Sm_Id:           newUpdate.Sm_Id,
		Name:            newUpdate.Name,
		SocialMedia_url: newUpdate.SocialMedia_url,
		User_id:         newUpdate.User_id,
		Updated_at:      newUpdate.Updated_at,
	}
	fmt.Println("ini photo handler : ", Update_respone)
	jsonData, _ := json.Marshal(Update_respone)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonData)

	fmt.Println("ini respone", Update_respone)
	return
}

func (s *SocialMediaHandler) SocialMediaDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if index, err := strconv.Atoi(id); err == nil {
		_, err := s.SocialMediaSrvc.SocialMediaDelete(r.Context(), id, &model.SocialMedia{
			Sm_Id: index,
		})
		if err != nil {
			w.Write([]byte(err.Error()))
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		msg := model.DeleteData{
			Message: "Your SocialMedia has been successfully deleted",
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(msg.Message))
		return
	}
}

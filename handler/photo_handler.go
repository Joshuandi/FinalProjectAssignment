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

type PhotoHandler struct {
	PhotoSrvc service.PhotoServiceInterface
}

func NewPhotohandler(PhotoSrvc service.PhotoServiceInterface) *PhotoHandler {
	return &PhotoHandler{PhotoSrvc: PhotoSrvc}
}

func (p *PhotoHandler) PhotoRegister(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := middleware.RunUser(ctx)
	fmt.Println("ini photo handler : ", userId.User_id)
	var photos *model.Photo
	//read dari body json
	jsonDec := json.NewDecoder(r.Body)
	err := jsonDec.Decode(&photos)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("ini photo handler : ", photos)
	//id := strconv.Itoa(userId.User_id)
	//throw to service
	newRegister, err := p.PhotoSrvc.PhotoRegister(r.Context(), &model.Photo{
		Title:     photos.Title,
		Caption:   photos.Caption,
		Photo_url: photos.Photo_url,
		User_id:   userId.User_id,
	})
	fmt.Println("ini photo handler : ", newRegister)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//thorw for respone result
	Register_respone := model.PhotoRegisterRespone{
		R_photo_id:   newRegister.Photo_id,
		R_title:      newRegister.Title,
		R_caption:    newRegister.Caption,
		R_photo_url:  newRegister.Photo_url,
		R_user_id:    newRegister.User_id,
		R_created_at: newRegister.Created_at,
	}
	fmt.Println("ini photo handler : ", Register_respone)
	jsonData, _ := json.Marshal(Register_respone)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonData)

	fmt.Println("ini respone", Register_respone)
	return
}

func (p *PhotoHandler) PhotoGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ids := params["id"]
	fmt.Println(ids)
	ctx := r.Context()
	userId := middleware.RunUser(ctx)
	fmt.Println("ini photo handler : ", userId.User_id)
	//var photos []*model.PhotoGet
	id := strconv.Itoa(userId.User_id)
	photo, err := p.PhotoSrvc.PhotoGet(ctx, id)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(&photo)
	w.Header().Add("Content-Type", "application/json")
	w.Write(res)
	return
}

func (p *PhotoHandler) PhotoUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var photos *model.Photo
	//read dari body json
	jsonDec := json.NewDecoder(r.Body)
	err := jsonDec.Decode(&photos)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("ini photo handler : ", photos)
	newUpdate, err := p.PhotoSrvc.PhotoUpdate(r.Context(), id, &model.Photo{
		Title:      photos.Title,
		Caption:    photos.Caption,
		Photo_url:  photos.Photo_url,
		Updated_at: photos.Updated_at,
	})
	fmt.Println("ini photo handler : ", newUpdate)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	Update_respone := model.PhotoUpdateRespone{
		U_photo_id:   newUpdate.Photo_id,
		U_title:      newUpdate.Title,
		U_caption:    newUpdate.Caption,
		U_photo_url:  newUpdate.Photo_url,
		U_user_id:    newUpdate.User_id,
		U_updated_at: newUpdate.Updated_at,
	}
	fmt.Println("ini photo handler : ", Update_respone)
	jsonData, _ := json.Marshal(Update_respone)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonData)

	fmt.Println("ini respone", Update_respone)
	return
}

func (p *PhotoHandler) PhotoDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if index, err := strconv.Atoi(id); err == nil {
		_, err := p.PhotoSrvc.PhotoDelete(r.Context(), id, &model.Photo{
			Photo_id: index,
		})
		if err != nil {
			w.Write([]byte(err.Error()))
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		msg := model.DeleteData{
			Message: "Your photo has been successfully deleted",
		}
		w.Write([]byte(msg.Message))
		return
	}
}

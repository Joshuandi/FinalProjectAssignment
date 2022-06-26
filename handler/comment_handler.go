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

type CommentHandler struct {
	CommentSrvc service.CommentServiceInterface
}

func NewCommenthandler(CommentSrvc service.CommentServiceInterface) *CommentHandler {
	return &CommentHandler{CommentSrvc: CommentSrvc}
}

func (c *CommentHandler) CommentRegister(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := middleware.RunUser(ctx)
	fmt.Println("ini Comment handler : ", userId.User_id)
	var comments *model.Comment
	//read dari body json
	jsonDec := json.NewDecoder(r.Body)
	err := jsonDec.Decode(&comments)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("ini Comment handler : ", comments)
	//id := strconv.Itoa(userId.User_id)
	//throw to service
	newRegister, err := c.CommentSrvc.CommentRegister(r.Context(), &model.Comment{
		Message:    comments.Message,
		Photo_id:   comments.Photo_id,
		User_id:    userId.User_id,
		Created_at: comments.Created_at,
	})
	fmt.Println("ini Comment handler : ", newRegister)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//thorw for respone result
	Register_respone := model.CommentRegisterRespone{
		R_Comment_id: newRegister.Comment_id,
		R_Message:    newRegister.Message,
		R_Photo_id:   newRegister.Photo_id,
		R_User_id:    newRegister.User_id,
		R_Created_at: newRegister.Created_at,
	}
	fmt.Println("ini Comment handler : ", Register_respone)
	jsonData, _ := json.Marshal(Register_respone)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(jsonData)

	fmt.Println("ini respone", Register_respone)
	return
}

func (c *CommentHandler) CommentGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ids := params["id"]
	fmt.Println(ids)

	ctx := r.Context()
	userId := middleware.RunUser(ctx)

	fmt.Println("ini comment handler : ", userId.User_id)
	//var photos []*model.PhotoGet
	id := strconv.Itoa(userId.User_id)
	comment, err := c.CommentSrvc.CommentGet(r.Context(), id)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(&comment)
	w.Header().Add("Content-Type", "application/json")
	w.Write(res)
	return
}

func (c *CommentHandler) CommentUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var cm *model.Comment
	//read dari body json
	jsonDec := json.NewDecoder(r.Body)
	err := jsonDec.Decode(&cm)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("ini photo handler : ", cm)
	newUpdate, err := c.CommentSrvc.CommentUpdate(r.Context(), id, &model.Comment{
		Message:    cm.Message,
		Updated_at: cm.Updated_at,
	})
	fmt.Println("ini photo handler : ", newUpdate)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	Update_respone := model.CommentShow{
		Comment_id: newUpdate.Comment_id,
		Title:      newUpdate.Title,
		Caption:    newUpdate.Caption,
		Photo_url:  newUpdate.Photo_url,
		User_id:    newUpdate.User_id,
		Updated_at: newUpdate.Updated_at,
	}
	fmt.Println("ini photo handler : ", Update_respone)
	jsonData, _ := json.Marshal(Update_respone)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonData)

	fmt.Println("ini respone", Update_respone)
	return
}

func (c *CommentHandler) CommentDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if index, err := strconv.Atoi(id); err == nil {
		_, err := c.CommentSrvc.CommentDelete(r.Context(), id, &model.Comment{
			Comment_id: index,
		})
		if err != nil {
			w.Write([]byte(err.Error()))
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		msg := model.DeleteData{
			Message: "Your comment has been successfully deleted",
		}
		w.Write([]byte(msg.Message))
		return
	}
}

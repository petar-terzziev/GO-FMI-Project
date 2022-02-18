package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"goproject/go-site-backend/DB"
	"goproject/go-site-backend/helpers"

	"github.com/gorilla/mux"
)

func PostThread(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't post thread."}
		json.NewEncoder(w).Encode(resp)
	}

	var formattedBody helpers.Post
	err = json.Unmarshal(body, &formattedBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't post thread."}
		json.NewEncoder(w).Encode(resp)
	}
	res, err := DB.PostThread(formattedBody.Title, formattedBody.Content, formattedBody.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't post thread."}
		json.NewEncoder(w).Encode(resp)
	}

	if len(res) > 0 {
		json.NewEncoder(w).Encode(formattedBody)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't post thread."}
		json.NewEncoder(w).Encode(resp)
	}
}

func PostComment(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't post comment."}
		json.NewEncoder(w).Encode(resp)
	}

	var formattedBody helpers.Comment
	err = json.Unmarshal(body, &formattedBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't post comment."}
		json.NewEncoder(w).Encode(resp)
	}
	res, err := DB.PostComment(formattedBody.Content, formattedBody.ThreadId, formattedBody.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't post comment."}
		json.NewEncoder(w).Encode(resp)
	}

	if len(res) > 0 {
		resp := helpers.ErrResponse{Message: res}
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := helpers.ErrResponse{Message: "Couldn't post comment."}
		json.NewEncoder(w).Encode(resp)
	}
}

func GetThreads(w http.ResponseWriter, r *http.Request) {
	res, err := DB.GetThreads()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode(res)
}

func GetComments(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	res, err := DB.GetComments(id["postId"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't get comments."}
		json.NewEncoder(w).Encode(resp)
	}
	json.NewEncoder(w).Encode(res)
}

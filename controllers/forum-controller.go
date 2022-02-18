package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"goproject/go-bank-backend/DB"
	"goproject/go-bank-backend/helpers"

	"github.com/gorilla/mux"
)

func PostThread(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	var formattedBody helpers.Post
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	res := DB.PostThread(formattedBody.Title, formattedBody.Content, formattedBody.UserId)

	if len(res) > 0 {
		resp := helpers.ErrResponse{Message: res}
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := helpers.ErrResponse{Message: "Couldn't post thread."}
		json.NewEncoder(w).Encode(resp)
	}
}

func PostComment(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	var formattedBody helpers.Comment
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	res := DB.PostComment(formattedBody.ThreadId, formattedBody.Content, formattedBody.UserId)

	if len(res) > 0 {
		resp := helpers.ErrResponse{Message: res}
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := helpers.ErrResponse{Message: "Couldn't post comment."}
		json.NewEncoder(w).Encode(resp)
	}
}

func GetThreads(w http.ResponseWriter, r *http.Request) {
	res := DB.GetThreads()
	json.NewEncoder(w).Encode(res)
}

func GetComments(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	res := DB.GetComments(id["postId"])
	json.NewEncoder(w).Encode(res)
}

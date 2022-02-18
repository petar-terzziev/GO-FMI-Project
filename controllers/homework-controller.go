package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"goproject/go-bank-backend/DB"
	"goproject/go-bank-backend/helpers"
)

func PostAssignment(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	var formattedBody helpers.Assignment
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	res := DB.PostAssignment(formattedBody.Title, formattedBody.Content, formattedBody.UserId)

	if len(res) > 0 {
		resp := helpers.ErrResponse{Message: res}
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := helpers.ErrResponse{Message: "Couldn't post assignment."}
		json.NewEncoder(w).Encode(resp)
	}
}

func PostSubmission(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	var formattedBody helpers.Submission
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	res := DB.PostComment(formattedBody.GithubLink, formattedBody.AssignmentId, formattedBody.UserId)

	if len(res) > 0 {
		resp := helpers.ErrResponse{Message: res}
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := helpers.ErrResponse{Message: "Couldn't register."}
		json.NewEncoder(w).Encode(resp)
	}
}

func GetAssignments(w http.ResponseWriter, r *http.Request) {
	res := DB.GetThreads()
	json.NewEncoder(w).Encode(res)
}

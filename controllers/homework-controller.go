package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"goproject/go-site-backend/DB"
	"goproject/go-site-backend/helpers"
)

func PostAssignment(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't post assignment."}
		json.NewEncoder(w).Encode(resp)
	}

	var formattedBody helpers.Assignment
	err = json.Unmarshal(body, &formattedBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't post assignment."}
		json.NewEncoder(w).Encode(resp)
	}
	res := DB.PostAssignment(formattedBody.Title, formattedBody.Content, formattedBody.UserId)

	if len(res) > 0 {
		json.NewEncoder(w).Encode(formattedBody)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't post assignment."}
		json.NewEncoder(w).Encode(resp)
	}
}

func PostSubmission(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't post assignment."}
		json.NewEncoder(w).Encode(resp)
	}

	var formattedBody helpers.Submission
	err = json.Unmarshal(body, &formattedBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't post assignment."}
		json.NewEncoder(w).Encode(resp)
	}
	res := DB.PostSubmission(formattedBody.GithubLink, formattedBody.AssignmentId, formattedBody.UserId)

	if len(res) > 0 {
		resp := helpers.ErrResponse{Message: res}
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := helpers.ErrResponse{Message: "Couldn't register."}
		json.NewEncoder(w).Encode(resp)
	}
}

func GetAssignments(w http.ResponseWriter, r *http.Request) {
	res := DB.GetAssignments()
	json.NewEncoder(w).Encode(res)
}

package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"goproject/go-site-backend/DB"
	"goproject/go-site-backend/helpers"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't login."}
		json.NewEncoder(w).Encode(resp)
	}

	var formattedBody helpers.LoginStruct
	err = json.Unmarshal(body, &formattedBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't login."}
		json.NewEncoder(w).Encode(resp)
	}
	login, err := DB.Login(formattedBody.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't login."}
		json.NewEncoder(w).Encode(resp)
	}

	if len(login.Username) > 0 {
		if err := bcrypt.CompareHashAndPassword([]byte(login.HashPassword), []byte(formattedBody.Password)); err != nil {
			resp := helpers.ErrResponse{Message: err.Error()}
			json.NewEncoder(w).Encode(resp)
		}
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    strconv.Itoa(int(login.ID)),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
		})

		token, err := claims.SignedString([]byte("MyBeautifulDarkTwistedFantasy"))
		if err == nil {
			json.NewEncoder(w).Encode(token)
		} else {
			resp := helpers.ErrResponse{Message: err.Error()}
			json.NewEncoder(w).Encode(resp)
		}
	} else {
		resp := helpers.ErrResponse{Message: "Wrong credemntials"}
		json.NewEncoder(w).Encode(resp)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't register."}
		json.NewEncoder(w).Encode(resp)
	}

	var formattedBody helpers.User
	err = json.Unmarshal(body, &formattedBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't register."}
		json.NewEncoder(w).Encode(resp)
	}
	password, err := bcrypt.GenerateFromPassword([]byte(formattedBody.Password), 14)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't register."}
		json.NewEncoder(w).Encode(resp)
	}
	register, err := DB.Register(formattedBody.Username, formattedBody.Email, string(password), formattedBody.FN)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't register."}
		json.NewEncoder(w).Encode(resp)
	}

	if len(register) > 0 {
		resp := helpers.ErrResponse{Message: register}
		json.NewEncoder(w).Encode(resp)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		resp := helpers.ErrResponse{Message: "Couldn't register."}
		json.NewEncoder(w).Encode(resp)
	}
}

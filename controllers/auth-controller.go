package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"goproject/go-bank-backend/DB"
	"goproject/go-bank-backend/helpers"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	var formattedBody helpers.LoginStruct
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := DB.Login(formattedBody.Username)

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
	helpers.HandleErr(err)

	var formattedBody helpers.LoginStruct
	err = json.Unmarshal(body, &formattedBody)
	password, _ := bcrypt.GenerateFromPassword([]byte(formattedBody.Password), 14)
	helpers.HandleErr(err)
	register := DB.Register(formattedBody.Username, string(password))

	if len(register) > 0 {
		resp := helpers.ErrResponse{Message: register}
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := helpers.ErrResponse{Message: "Couldn't register."}
		json.NewEncoder(w).Encode(resp)
	}
}

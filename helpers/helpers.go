package helpers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Post struct {
	ID      uint
	Title   string
	Content string
	Poster  string
	UserId  string
}

type User struct {
	Username string
	Email    string
	Password string
	FN       string
}

type Comment struct {
	ID       uint
	Title    string
	Content  string
	ThreadId string
	UserId   string
	Poster   string
}

type Assignment struct {
	ID      int
	Title   string
	Content string
	UserId  string
}

type Submission struct {
	ID           int
	GithubLink   string
	AssignmentId string
	UserId       string
}

type LoginStruct struct {
	Email    string
	Password string
}

type Response struct {
	Data string
}

type ErrResponse struct {
	Message string
}

func HandleErr(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}

func HashAndSalt(pass []byte) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)

	return string(hashed), err
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return "MyBeautifulDarkTwistedFantasy", nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

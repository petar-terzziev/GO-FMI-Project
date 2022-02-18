package api

import (
	"fmt"
	"log"
	"net/http"

	"goproject/go-bank-backend/controllers"

	"github.com/gorilla/mux"
)

type WithCORS struct {
	r *mux.Router
}

func (s *WithCORS) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		res.Header().Set("Access-Control-Allow-Origin", origin)
		res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		res.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		res.Header().Set("Access-Control-Allow-Credentials", "true")
	}

	// Stop here for a Preflighted OPTIONS request.
	if req.Method == "OPTIONS" {
		return
	}
	log.Printf("not options")
	log.Print(res.Header().Get("Access-Control-Allow-Origin"))
	// Let Gorilla work
	s.r.ServeHTTP(res, req)
}

func StartApi() {

	router := mux.NewRouter()
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/forum/postThread", controllers.PostThread).Methods("POST")
	router.HandleFunc("/forum/postComment", controllers.PostComment).Methods("POST")
	router.HandleFunc("/forum/getComments/{postId}", controllers.PostThread).Methods("GET")
	router.HandleFunc("/forum/getAll", controllers.GetThreads).Methods("GET")

	fmt.Println("App is working on port :8888")
	http.Handle("/", &WithCORS{router})
	log.Fatal(http.ListenAndServe(":8888", nil))

}

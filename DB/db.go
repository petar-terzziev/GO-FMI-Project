package DB

import (
	"database/sql"

	"goproject/go-bank-backend/helpers"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/lib/pq"
)

type User struct {
	ID           int
	Username     string
	HashPassword string
}

type Account struct {
	ID      int
	Name    string
	Balance int
}

func connectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/GO_APP?parseTime=true")
	helpers.HandleErr(err)
	return db

}

// func correctAuthUser(accPass string, typedPass string) bool {
// 	accPassByte := []byte(accPass)
// 	typedPassByte := []byte(typedPass)
// 	err := bcrypt.CompareHashAndPassword(accPassByte, typedPassByte)
// 	if err != nil {
// 			return false
// 	}
//   return true
// }

func Login(username string) User {
	db := connectDB()
	results, err := db.Query("SELECT id, username, password FROM users x WHERE username=?", username)
	helpers.HandleErr(err)
	var users []User

	for results.Next() {
		var user User
		err := results.Scan(&user.ID, &user.Username, &user.HashPassword)
		helpers.HandleErr(err)
		users = append(users, user)
	}
	defer results.Close()
	return users[0]
}

func Register(username string, pass string) string {
	db := connectDB()
	results, err := db.Query("INSERT INTO users (username, password) Values (?, ?)", username, pass)
	helpers.HandleErr(err)
	defer results.Close()
	return "success"

}

func PostThread(title string, content string, userId string) string {
	// call, err := db.Query(dbCall("SELECT id, username, email FROM users WHERE username= ? AND password= ?")
	db := connectDB()
	results, err := db.Query("INSERT INTO threads (title, content, user_id) Values ( ?, ?, ?)", title, content, userId)
	helpers.HandleErr(err)
	defer results.Close()
	return "success"
}

func PostComment(content string, threadId string, userId string) string {
	db := connectDB()
	results, err := db.Query("INSERT INTO comments (content, thread_id, user_id) Values (?, ?, ?)", content, threadId, userId)
	helpers.HandleErr(err)
	defer results.Close()
	return "success"
}

func PostAssignment(title string, content string, userId string) string {
	db := connectDB()
	results, err := db.Query("INSERT INTO homework_assignments (title, content, user_id) Values (?, ?, ?)", title, content, userId)
	helpers.HandleErr(err)
	defer results.Close()
	return "success"
}

func GetAssignments() []helpers.Assignment {
	db := connectDB()
	results, err := db.Query("SELECT id, title, content, user_id FROM homework_assignments")
	helpers.HandleErr(err)
	var homeworks []helpers.Assignment

	for results.Next() {
		var hw helpers.Assignment
		err := results.Scan(&hw.ID, &hw.Title, &hw.Content, &hw.UserId)
		helpers.HandleErr(err)
		homeworks = append(homeworks, hw)
	}
	defer results.Close()
	return homeworks
}
func PostSubmission(githubLink string, assignmentId string, userId string) string {
	db := connectDB()
	results, err := db.Query("INSERT INTO homework_submissions ( github_link, assignment_id, user_id) Values (?, ?, ?)", githubLink, assignmentId, userId)
	helpers.HandleErr(err)
	defer results.Close()
	return "success"
}

func GetSubmissionsByAssignment(assignmentId string) []helpers.Submission {
	db := connectDB()
	results, err := db.Query("SELECT id, github_link, user_id FROM homework_submissions WHERE assignment_id =?", assignmentId)
	helpers.HandleErr(err)
	var submissions []helpers.Submission

	for results.Next() {
		var submission helpers.Submission
		submission.AssignmentId = assignmentId
		err := results.Scan(&submission.ID, &submission.GithubLink, &submission.UserId)
		helpers.HandleErr(err)
		submissions = append(submissions, submission)
	}
	defer results.Close()
	return submissions
}
func GetSubmissionsByUser(userId string) []helpers.Submission {
	db := connectDB()
	results, err := db.Query("SELECT id, github_link, user_id FROM homework_submissions WHERE user_id =?", userId)
	helpers.HandleErr(err)
	var submissions []helpers.Submission

	for results.Next() {
		var submission helpers.Submission
		submission.UserId = userId
		err := results.Scan(&submission.ID, &submission.GithubLink, &submission.AssignmentId)
		helpers.HandleErr(err)
		submissions = append(submissions, submission)
	}
	defer results.Close()
	return submissions
}

func GetThreads() []helpers.Post {
	db := connectDB()
	results, err := db.Query("SELECT id, title, content, user_id FROM threads")
	helpers.HandleErr(err)
	var posts []helpers.Post

	for results.Next() {
		var post helpers.Post
		err := results.Scan(&post.ID, &post.Title, &post.Content, &post.UserId)
		helpers.HandleErr(err)
		posts = append(posts, post)
	}
	defer results.Close()
	return posts
}

func GetComments(ThreadId string) []helpers.Comment {
	db := connectDB()
	results, err := db.Query("SELECT id, content, userId FROM threads WHERE threadId = ?", ThreadId)
	helpers.HandleErr(err)
	var posts []helpers.Comment

	for results.Next() {
		var post helpers.Comment
		err := results.Scan(&post.ID, &post.Content, &post.UserId)
		helpers.HandleErr(err)
		posts = append(posts, post)
	}
	defer results.Close()
	return posts
}

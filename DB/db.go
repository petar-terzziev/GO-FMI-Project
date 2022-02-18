package DB

import (
	"database/sql"

	"goproject/go-site-backend/helpers"

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

//Gets a user from the db
func Login(username string) (User, error) {
	db := connectDB()
	results, err := db.Query("SELECT id, username, password FROM users x WHERE email=?", username)
	if err != nil {
		defer results.Close()
		return User{}, err
	}
	var users []User

	for results.Next() {
		var user User
		err := results.Scan(&user.ID, &user.Username, &user.HashPassword)
		if err != nil {
			defer results.Close()
			return User{}, err
		}
		users = append(users, user)
	}
	if err != nil {
		defer results.Close()
		return User{}, err
	} else {
		if len(users) > 0 {
			defer results.Close()
			return users[0], nil
		} else {
			defer results.Close()
			return User{}, err
		}
	}
}

// Inserts a user into the database users table
func Register(username string, email string, pass string, fn string) (string, error) {
	db := connectDB()
	results, err := db.Query("INSERT INTO users (username, email, password, fn) Values (?, ?, ?, ?)", username, email, pass, fn)
	if err != nil {
		defer results.Close()
		return "failure", err
	} else {
		defer results.Close()
		return "success", nil
	}

}

//Inserts a thread into the threads table
func PostThread(title string, content string, userId string) (string, error) {
	db := connectDB()
	results, err := db.Query("INSERT INTO threads (title, content, user_id) Values ( ?, ?, ?)", title, content, userId)
	if err != nil {
		defer results.Close()
		return "failure", err
	} else {
		defer results.Close()
		return "success", nil
	}
}

func PostComment(content string, threadId string, userId string) (string, error) {
	db := connectDB()
	results, err := db.Query("INSERT INTO comments (content, thread_id, user_id) Values (?, ?, ?)", content, threadId, userId)
	if err != nil {
		defer results.Close()
		return "failure", err
	} else {
		defer results.Close()
		return "success", nil
	}
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

func GetThreads() ([]helpers.Post, error) {
	db := connectDB()
	results, err := db.Query("SELECT threads.id, title, content, COALESCE(username,'notexisting') FROM threads LEFT JOIN users ON user_id = users.id")
	if err != nil {
		defer results.Close()
		return []helpers.Post{}, err
	}
	var posts []helpers.Post

	for results.Next() {
		var post helpers.Post
		err := results.Scan(&post.ID, &post.Title, &post.Content, &post.Poster)
		if err != nil {
			defer results.Close()
			return posts, err
		}
		posts = append(posts, post)
	}

	defer results.Close()
	return posts, nil

}

func GetComments(ThreadId string) ([]helpers.Comment, error) {
	db := connectDB()
	results, err := db.Query("SELECT comments.id, content, COALESCE(username,'notexisting') FROM comments LEFT JOIN users on user_id = users.id WHERE thread_id = ?", ThreadId)
	if err != nil {
		defer results.Close()
		return []helpers.Comment{}, err
	}
	var posts []helpers.Comment

	for results.Next() {
		var post helpers.Comment
		err := results.Scan(&post.ID, &post.Content, &post.Poster)
		if err != nil {
			defer results.Close()
			return posts, err
		}
		posts = append(posts, post)
	}
	defer results.Close()
	return posts, nil
}

package migrations

import (
	"fmt"

	"goproject/go-site-backend/helpers"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Thread struct {
	gorm.Model
	Title   string
	Content string `gorm:"type:text"`
	UserID  uint
	User    User `gorm:"foreignKey:UserID"`
}

type Comment struct {
	gorm.Model
	Content  string `gorm:"type:text"`
	ThreadID uint
	Thread   Thread `gorm:"foreignKey:ThreadID"`
	UserID   uint
	User     User `gorm:"foreignKey:UserID"`
}

type HomeworkAssignment struct {
	gorm.Model
	Title   string
	Content string
	UserId  int
	User    User `gorm:"foreignKey:UserID"`
}

type HomeworkSubmssion struct {
	gorm.Model
	GithubLink   string
	AssignmentId int
	Assignment   HomeworkAssignment `gorm:"foreignKey:AssignmentID"`
	UserId       int
	User         User `gorm:"foreignKey:UserID"`
}

func connectDB() *gorm.DB {
	// dsn := "opds:opds@tcp(127.0.0.1:3306)/GO_APP?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open("mysql", "root:1234@tcp(localhost:3306)/GO_APP?parseTime=true")
	helpers.HandleErr(err)
	return db
}

// This is correct way of creating password
// func HashAndSalt(pass []byte) string {
// 	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
// 	helpers.HandleErr(err)
// 	return string(hashed)
// }

func Migrate() {
	db := connectDB()
	fmt.Println('y')
	db.AutoMigrate(&User{}, &Thread{}, &Comment{}, &HomeworkAssignment{}, &HomeworkSubmssion{})
	defer db.Close()
}

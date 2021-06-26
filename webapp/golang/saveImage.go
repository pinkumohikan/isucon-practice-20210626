package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)
type Post21 struct {
	ID           int       `db:"id"`
	UserID       int       `db:"user_id"`
	Imgdata      []byte    `db:"imgdata"`
	Body         string    `db:"body"`
	Mime         string    `db:"mime"`
	CreatedAt    time.Time `db:"created_at"`
	CommentCount int
	CSRFToken    string
}
func main() {
	host := "isucon-app3"
	port := "3306"
	user := "isuconp"
	password := "isuconp"
	dbname := "isuconp"

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %s.", err.Error())
	}
	defer db.Close()

	post := Post21{}
	err = db.Get(&post, "SELECT * FROM `posts`")
	if err != nil {
		log.Print(err)
		return
	}
	imageType :=strings.Replace(post.Mime,"image", "", 1)
	file, _ := os.Create("./images/" + strconv.Itoa(post.ID) + "." + imageType)
	fmt.Println("./images/" + strconv.Itoa(post.ID) + "." + imageType)
	defer file.Close()
	file.Write(([]byte)(post.Imgdata))
	fmt.Println("testtest")
}

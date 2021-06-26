package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

	posts := []Post21{}
	err = db.Select(&posts, "SELECT * FROM `posts`")
	if err != nil {
		log.Print(err)
		return
	}
	for _, p := range posts {
		imageType :=strings.Replace(p.Mime,"image/", "", 1)
		file, _ := os.Create("/home/isucon/isucon-practice-20210626/webapp/golang/images/" + strconv.Itoa(p.ID) + "." + imageType)
		fmt.Println("./images/" + strconv.Itoa(p.ID) + "." + imageType)
		defer file.Close()
		file.Write(([]byte)(p.Imgdata))
		fmt.Println("testtest")
	}

}

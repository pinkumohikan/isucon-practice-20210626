package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	host := os.Getenv("ISUCONP_DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("ISUCONP_DB_PORT")
	if port == "" {
		port = "3306"
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Failed to read DB port number from an environment variable ISUCONP_DB_PORT.\nError: %s", err.Error())
	}
	user := os.Getenv("ISUCONP_DB_USER")
	if user == "" {
		user = "root"
	}
	password := os.Getenv("ISUCONP_DB_PASSWORD")
	dbname := os.Getenv("ISUCONP_DB_NAME")
	if dbname == "" {
		dbname = "isuconp"
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %s.", err.Error())
	}
	defer db.Close()

	post := Post{}
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

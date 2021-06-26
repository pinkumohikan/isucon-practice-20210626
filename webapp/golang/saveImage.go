package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	post := Post{}
	err := db.Get(&post, "SELECT * FROM `posts`")
	if err != nil {
		log.Print(err)
		return
	}
	imageType :=strings.Replace(post.Mime,"image", "", 1)
	file, _ := os.Create(`./images/` + strconv.Itoa(post.ID) + "." + imageType)
	defer file.Close()
	file.Write(([]byte)(post.Imgdata))
	fmt.Println("testtest")
}

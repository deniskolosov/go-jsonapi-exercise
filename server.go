package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

type Post struct {
	Id        int        `jsonapi:"primary,posts"`
	Title     string     `jsonapi:"attr,title"`
	Content   string     `jsonapi:"attr,content"`
	Comments  []*Comment `jsonapi:"relation,comments"`
	Tags      []*Tag     `jsonapi:"relation,tags"`
	CreatedAt int        `jsonapi:"attr,created_at"`
	UpdatedAt int        `jsonapi:"attr,updated_at"`
}

type Comment struct {
	Id        int    `jsonapi:"primary,comments"`
	Content   string `jsonapi:"attr,content"`
	PostId    int    `jsonapi:"attr,post_id"`
	CreatedAt int    `jsonapi:"attr,created_at"`
	UpdatedAt int    `jsonapi:"attr,updated_at"`
}

type Tag struct {
	Id        int     `jsonapi:"primary,tag"`
	Name      string  `jsonapi:"attr,name"`
	CreatedAt int     `jsonapi:"attr,created_at"`
	Posts     []*Post `jsonapi:"relation,posts"`
}

type PostsTags struct {
	Id     int `jsonapi:"primary,posts_tags"`
	PostId int `jsonapi:"attr,post_id"`
	TagId  int `jsonapi:"attr,tag_id"`
}

var db *sql.DB

func initDB(dataSourceName string) {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	onError(err)
}

func main() {
	initDB("./database.db")
	defer db.Close()

	var port string
	flag.StringVar(&port, "p", ":8080", "Port on which service will run")
	flag.Parse()

	http.HandleFunc("/posts", PostsHandler)
	http.HandleFunc("/comments", CommentsHandler)
	http.HandleFunc("/tags", TagsHandler)
	fmt.Println("Running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))

}

func onError(err error) {
	if err != nil {
		panic(err)
	}
}

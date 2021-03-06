package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Post struct {
	ID      int
	Content string
	Author  string
}

var Db *sql.DB

func main() {
	post := Post{}
	post.Content = "content from go"
	post.Author = "author from go"
	fmt.Println(post)
	fmt.Println("Hello world!!!")

	// post.create()
	// fmt.Println(post)
}

// connect to the Db
func init() {
	var err error
	Db, err = sql.Open("postgres", "user=grs password=grs dbname=grs_db sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// Get the list of post
func list() (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts")

	for i := 0; rows.Next(); i = i + 1 {
		posts = append(posts, Post{})
		if err := rows.Scan(&posts[i].ID, &posts[i].Content, &posts[i].Author); err != nil {
			fmt.Println(err)
		}
	}
	return
}

// Get a single post
func retrieve(id int) (post Post, err error) {
	// post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.ID, &post.Content, &post.Author)
	return
}

// Create a new post
func (post *Post) create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.ID)
	return
}

// Update a post
func (post *Post) update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.ID, post.Content, post.Author)
	return
}

// Delete a post
func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.ID)
	return
}

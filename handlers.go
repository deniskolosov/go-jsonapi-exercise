package main

import (
	"github.com/google/jsonapi"
	"net/http"
	"regexp"
)

func getPosts() ([]interface{}, error) {
	postRows, err := db.Query("select * from post;")
	onError(err)
	defer postRows.Close()
	var posts []*Post

	for postRows.Next() {
		var p Post
		onError(postRows.Scan(&p.Id, &p.Title, &p.Content, &p.CreatedAt, &p.UpdatedAt))
		commentsRows, err := db.Query("select * from comment where post_id=?", p.Id)
		onError(err)

		for commentsRows.Next() {
			var c Comment
			onError(commentsRows.Scan(&c.Id, &c.Content, &c.PostId, &c.CreatedAt, &c.UpdatedAt))
			p.Comments = append(p.Comments, &c)
		}
		tagsRows, err := db.Query("select posts_tags.tag_id, tag.name,"+
			" tag.created_at from posts_tags join tag on posts_tags.tag_id = tag.id  where post_id=?;", p.Id)
		onError(err)

		for tagsRows.Next() {
			var t Tag
			onError(tagsRows.Scan(&t.Id, &t.Name, &t.CreatedAt))
			p.Tags = append(p.Tags, &t)
		}

		posts = append(posts, &p)

	}
	postsInterface := make([]interface{}, len(posts))
	for i, post := range posts {
		postsInterface[i] = post
	}
	return postsInterface, nil
}

func getComments() ([]interface{}, error) {
	commentRows, err := db.Query("select * from comment;")
	onError(err)
	defer commentRows.Close()
	var comments []*Comment
	for commentRows.Next() {
		var c Comment
		onError(commentRows.Scan(&c.Id, &c.Content, &c.PostId, &c.CreatedAt, &c.UpdatedAt))

		comments = append(comments, &c)

	}
	commentsInterface := make([]interface{}, len(comments))
	for i, comment := range comments {
		commentsInterface[i] = comment
	}
	return commentsInterface, nil
}

func getTags() ([]interface{}, error) {
	tagRows, err := db.Query("select * from tag;")
	onError(err)
	defer tagRows.Close()
	var tags []*Tag
	for tagRows.Next() {
		var t Tag
		onError(tagRows.Scan(&t.Id, &t.Name, &t.CreatedAt))

		postRows, err := db.Query("select post.id, post.title, post.content, post.created_at, post.updated_at"+
			" from post"+
			" join posts_tags"+
			" on post.id = posts_tags.post_id WHERE tag_id=?;", t.Id)
		onError(err)

		for postRows.Next() {
			var p Post
			onError(postRows.Scan(&p.Id, &p.Title, &p.Content, &p.CreatedAt, &p.UpdatedAt))
			t.Posts = append(t.Posts, &p)
		}

		tags = append(tags, &t)
	}
	tagsInterface := make([]interface{}, len(tags))
	for i, post := range tags {
		tagsInterface[i] = post
	}
	return tagsInterface, nil
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, _ := getPosts()
	handlerFactory(posts)(w, r)
}

func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	comments, _ := getComments()
	handlerFactory(comments)(w, r)
}

func TagsHandler(w http.ResponseWriter, r *http.Request) {
	tags, _ := getTags()
	handlerFactory(tags)(w, r)
}

func handlerFactory(entities []interface{}) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !regexp.MustCompile(`application/vnd\.api\+json`).Match([]byte(r.Header.Get("Accept"))) {
			http.Error(w, "Unsupported Media Type", 415)
			return
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/vnd.api+json")
		if err := jsonapi.MarshalManyPayload(w, entities); err != nil {
			http.Error(w, err.Error(), 500)
		}
	}
}

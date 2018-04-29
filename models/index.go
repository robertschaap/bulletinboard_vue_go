package models

import (
	"log"
	"upper.io/db.v3/mongo"
)

// Database Settings
var settings = mongo.ConnectionURL{
  Host:       "localhost:27017",
  Database:   "bulletinboard",
}

// Comment struct
type Comment struct {
	Title		string	`json:"title"`
	Body		string	`json:"body"`
	Avatar	string	`json:"avatar"`
	Name		string	`json:"name"`
}

// GetComments sends data to controler
func GetComments() []Comment {
	var sess, err = mongo.Open(settings)
	if err != nil {
		log.Fatalf("Cannot open: %p", err)
	}
	defer sess.Close()

	res := sess.Collection("comments").Find()

	var comments []Comment
	if err := res.All(&comments); err != nil {
		log.Fatalf("Cannot read comments: %p", err)
	}

	for _, comment := range comments {
		log.Printf("%q\n", comment.Title)
	}

	return comments
}

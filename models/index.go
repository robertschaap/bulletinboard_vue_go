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
func GetComments(page uint, sort string) []Comment {
	var sess, err = mongo.Open(settings)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	res := sess.Collection("comments").
		Find().
		OrderBy("createdAt DESC").
		Paginate(4)

	var comments []Comment

	err = res.Page(page).All(&comments)

	if err != nil {
		log.Fatal(err)
	}

	return comments
}

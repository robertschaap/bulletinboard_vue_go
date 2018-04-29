package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// Comment struct
type Comment struct {
	Title		string	`json:"title"`
	Body		string	`json:"body"`
	Avatar	string	`json:"avatar"`
	Name		string	`json:"name"`
}

func getComments(w http.ResponseWriter, r *http.Request) {
	var comments []Comment
	comments = append(comments, Comment{Title: "Test", Body: "Test", Avatar: "bunny", Name: "Test" })

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

func createComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/comment", getComments).Methods("GET")
	router.HandleFunc("/api/comment/new", createComment).Methods("POST")

	port := ":3000"
	fmt.Println("Running on port"+port)
	log.Fatal(http.ListenAndServe(port, router))
}

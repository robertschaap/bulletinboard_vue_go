package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	"../models"
)

func getCommentHandler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	sort := v.Get("sort")
	offset, _ := strconv.ParseUint(v.Get("offset"), 10, 32)

	comment := models.GetComments(uint(offset), sort)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comment)
}

func postCommentHandler(w http.ResponseWriter, r *http.Request) {
	var formbody models.Comment
	json.NewDecoder(r.Body).Decode(&formbody)

	models.PostComment(formbody)
}

// APIController Routes
func APIController(port string) {

	router := mux.NewRouter()
	router.HandleFunc("/api/comment", getCommentHandler).Methods("GET")
	router.HandleFunc("/api/comment/new", postCommentHandler).Methods("POST")

	fmt.Println("Running on port"+port)
	log.Fatal(http.ListenAndServe(port, router))
}

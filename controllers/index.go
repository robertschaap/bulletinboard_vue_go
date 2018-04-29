package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"../models"
)


func getComments(w http.ResponseWriter, r *http.Request) {
	comment := models.GetComments()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comment)
}

// APIController Routes
func APIController(p string) {

	router := mux.NewRouter()
	router.HandleFunc("/api/comment", getComments).Methods("GET")

	fmt.Println("Running on port"+p)
	log.Fatal(http.ListenAndServe(p, router))
}

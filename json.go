package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sqsiek0/go_learning_project/models"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Błąd podczas parsowania:", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Problem z serwerem")
	}
	type errorMessage struct {
		ErrorMessage string `json:"error"`
	}

	respondWithJSON(w, code, errorMessage{ErrorMessage: msg})
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}

	users := []models.UserModel{
		{ID: 1, Name: "Maks"},
		{ID: 2, Name: "Kacper"},
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Printf("Error encoding users: %v", err)
		http.Error(w, "Error encoding users", http.StatusInternalServerError)
		return
	}

}

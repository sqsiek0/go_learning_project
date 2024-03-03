package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sqsiek0/go_learning_project/internal/database"
)

func (apiCnf *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name    string `json:"name"`
		Surname string `json:"surname"`
	}
	params := parameters{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintln("Error parsing JSON", err))
		return
	}

	user, err := apiCnf.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Createdat: time.Now().UTC(),
		Updatedat: time.Now().UTC(),
		Name:      params.Name,
		Surname:   params.Surname,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintln("Can't create user", err))
		return
	}

	respondWithJSON(w, 200, changeUserTitles(user))
}
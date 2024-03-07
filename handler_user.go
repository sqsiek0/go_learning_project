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
	if params.Name == "" {
		respondWithError(w, 400, fmt.Sprintln("Invalid name provided"))
		return
	} else if params.Surname == "" {
		respondWithError(w, 400, fmt.Sprintln("Invalid surname provided"))
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

	listOfUsers := make([]database.User, 0)
	respondWithJSON(w, 200, changeUserTitles(append(listOfUsers, user)))
}

func (apiCnf *apiConfig) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := apiCnf.DB.GetUsers(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintln("Error during getting users", err))
	}

	respondWithJSON(w, 200, changeUserTitles(users))
}

func (apiCnf *apiConfig) handleGetUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {
	listOfUsers := make([]database.User, 0)
	respondWithJSON(w, 200, changeUserTitles(append(listOfUsers, user)))
}

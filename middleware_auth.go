package main

import (
	"fmt"
	"net/http"

	"github.com/sqsiek0/go_learning_project/internal/auth"
	"github.com/sqsiek0/go_learning_project/internal/database"
)

func (apiCnf *apiConfig) middlewareAuth(handler func(http.ResponseWriter, *http.Request, database.User)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		value, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Invalid headers"))
			return
		}

		user, err := apiCnf.DB.GetUserByApiKey(r.Context(), value)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Error during searching for user"))
			return
		}

		handler(w, r, user)
	}
}

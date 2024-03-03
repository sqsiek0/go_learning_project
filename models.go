package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/sqsiek0/go_learning_project/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Createdat time.Time `json:"created_at"`
	Updatedat time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
}

func changeUserTitles(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Createdat: dbUser.Createdat,
		Updatedat: dbUser.Updatedat,
		Name:      dbUser.Name,
		Surname:   dbUser.Surname,
	}
}

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

func changeUserTitles(dbUser []database.User) []User {
	listOfUsers := make([]User, 0)

	for _, v := range dbUser {
		listOfUsers = append(listOfUsers, User{
			ID:        v.ID,
			Createdat: v.Createdat,
			Updatedat: v.Updatedat,
			Name:      v.Name,
			Surname:   v.Surname,
		})
	}
	return listOfUsers

	// return User{
	// 	ID:        dbUser.ID,
	// 	Createdat: dbUser.Createdat,
	// 	Updatedat: dbUser.Updatedat,
	// 	Name:      dbUser.Name,
	// 	Surname:   dbUser.Surname,
	// }
}

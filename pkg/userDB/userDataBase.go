package userDataBaseP

import "github.com/google/uuid"

type User struct {
	Age     int       `json:"age"`
	Name    string    `json:"name"`
	Id      uuid.UUID `json:"id"`
	Friends []string  `json:"friends"`
}

var UserDataBaseP map[uuid.UUID]User

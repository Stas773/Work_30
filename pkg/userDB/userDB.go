package userDB

import "github.com/google/uuid"

type User struct {
	Age     int       `json:"age"`
	Name    string    `json:"name"`
	Id      uuid.UUID `json:"id"`
	Friends []string  `json:"friends"`
}

var UserDB map[uuid.UUID]User

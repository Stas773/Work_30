package main

import (
	"net/http"

	"D/Works/GO/30work/pkg/createUser"
	"D/Works/GO/30work/pkg/deleteUser"
	"D/Works/GO/30work/pkg/getAllUsers"
	"D/Works/GO/30work/pkg/getUser"
	"D/Works/GO/30work/pkg/getUserFriends"
	"D/Works/GO/30work/pkg/makeFriends"
	"D/Works/GO/30work/pkg/updateUser"
	"D/Works/GO/30work/pkg/userDB"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func main() {
	router := chi.NewRouter()

	userDB.UserDB = make(map[uuid.UUID]userDB.User)

	router.Post("/create", createUser.CreateUser)
	router.Post("/make_friends", makeFriends.MakeFriends)
	router.Get("/Get/{userId}", getUser.GetUser)
	router.Get("/Get/all", getAllUsers.GetAllUsers)
	router.Get("/Get/friends/{userId}", getUserFriends.GetUserFriends)
	router.Patch("/Patch/{userId}", updateUser.UpdateUser)
	router.Delete("/Delete/{userId}", deleteUser.DeleteUser)

	http.ListenAndServe(":8080", router)
}

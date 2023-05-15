package createUser

import (
	"D/Works/GO/30work/pkg/userDB"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user userDB.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Id = uuid.New()

	userDB.UserDB[user.Id] = user

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User was created"))
}

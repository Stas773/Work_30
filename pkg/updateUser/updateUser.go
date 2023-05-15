package updateUser

import (
	"D/Works/GO/30work/pkg/userDB"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")

	userIdInt, err := uuid.Parse(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, ok := userDB.UserDB[userIdInt]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var updateUser userDB.User
	err = json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if updateUser.Age > 0 {
		user.Age = updateUser.Age
	}
	if updateUser.Name != "" {
		user.Name = updateUser.Name
	}
	userDB.UserDB[user.Id] = user

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %s updated", userIdInt)
}

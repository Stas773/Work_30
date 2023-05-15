package makeFriends

import (
	"D/Works/GO/30work/pkg/userDB"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func MakeFriends(w http.ResponseWriter, r *http.Request) {
	type FriendRequest struct {
		User1 uuid.UUID `json:"user1"`
		User2 uuid.UUID `json:"user2"`
	}

	var friendRequest FriendRequest

	err := json.NewDecoder(r.Body).Decode(&friendRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user1, ok := userDB.UserDB[friendRequest.User1]
	if !ok {
		http.Error(w, "User1 not found", http.StatusNotFound)
		return
	}

	user2, ok := userDB.UserDB[friendRequest.User2]
	if !ok {
		http.Error(w, "User2 not found", http.StatusNotFound)
		return
	}

	user1.Friends = append(user1.Friends, user2.Id.String())
	user2.Friends = append(user2.Friends, user1.Id.String())

	userDB.UserDB[user1.Id] = user1
	userDB.UserDB[user2.Id] = user2

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User %s and %s are friends now", user1.Id, user2.Id)
}

package deleteUser

import (
	"D/Works/GO/30work/pkg/userDB"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
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

	for _, friendIdStr := range user.Friends {
		friendId, err := uuid.Parse(friendIdStr)
		if err != nil {
			continue
		}

		friend, ok := userDB.UserDB[friendId]
		if !ok {
			continue
		}

		for i, id := range friend.Friends {
			if id == userId {
				friend.Friends = append(friend.Friends[:i], friend.Friends[i+1:]...)
			}
		}
		userDB.UserDB[friend.Id] = friend
	}

	delete(userDB.UserDB, userIdInt)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %s deleted", userIdInt)
}

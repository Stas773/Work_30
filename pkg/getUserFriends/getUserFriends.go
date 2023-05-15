package getUserFriends

import (
	"D/Works/GO/30work/pkg/userDB"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetUserFriends(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "userId")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid Id", http.StatusBadRequest)
		return
	}
	user, ok := userDB.UserDB[id]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user.Friends)
}

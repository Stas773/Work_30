package getAllUsers

import (
	"D/Works/GO/30work/pkg/userDB"
	"encoding/json"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	allUsers := userDB.UserDB
	if len(allUsers) == 0 {
		http.Error(w, "Users not found", http.StatusNotFound)
		return
	}
	for _, v := range allUsers {
		json.NewEncoder(w).Encode(v)
	}
}

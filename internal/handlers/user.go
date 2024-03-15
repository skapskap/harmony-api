package handlers

import (
	"encoding/json"
	"net/http"
)

func UserInfo(w http.ResponseWriter, r *http.Request) {
	userInfo := map[string]string{
		"username": "Gabriel Caldas",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userInfo)
}

package handlers

import (
	"encoding/json"
	"net/http"
)

func userInfo(w http.ResponseWriter, r *http.Request) {
	userInfo := map[string]string{
		"username": "Gabriel Caldas",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userInfo)
}
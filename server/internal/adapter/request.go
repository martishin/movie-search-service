package adapter

import (
	"net/http"
	"strconv"

	"github.com/markbates/goth/gothic"
)

func GetUserIDFromSession(r *http.Request) (int, error) {
	// Retrieve user ID from session
	userIDStr, err := gothic.GetFromSession("user_id", r)
	if err != nil || userIDStr == "" {
		return 0, err
	}

	// Convert userID from string to int
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

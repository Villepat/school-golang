package session

import (
	"net/http"
)

func GetUserName(w http.ResponseWriter, r *http.Request) (string, error) {
	if validateSession(w, r) {
		c, err := r.Cookie("session_token")
		if err != nil {
			return "", err
		}
		sessionToken := c.Value
		userSession := sessions[sessionToken]
		//RefreshSesh(w, r)
		return userSession.username, nil
	}
	return "", nil
}

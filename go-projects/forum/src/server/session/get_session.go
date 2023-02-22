package session

import (
	"net/http"
	"time"
)

func validateSession(w http.ResponseWriter, r *http.Request) bool {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			return false
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	sessionToken := c.Value
	// We then get the session from our session map
	userSession, exists := sessions[sessionToken]
	if !exists {
		// If the session token is not present in session map, return an unauthorized error
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   "",
			Expires: time.Now(),
		})
		return false
	}
	// If the session is present, but has expired, we can delete the session, and return
	// an unauthorized status
	if userSession.isExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	return true
}

func GetSession(w http.ResponseWriter, r *http.Request) (bool, string) {
	if validateSession(w, r) {
		c, _ := r.Cookie("session_token")
		sessionToken := c.Value
		userSession := sessions[sessionToken]
		//RefreshSesh(w, r)
		return true, userSession.username
	}
	return false, ""
}

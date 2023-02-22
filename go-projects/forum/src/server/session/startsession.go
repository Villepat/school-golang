package session

import (
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

// Execute this function when a submit with valid credentials is made on the Login page.
// Creates an UUID session token with an expiration date for the user.
func StartSesh(w http.ResponseWriter, r *http.Request) {
	for k, v := range sessions {
		if v.username == r.FormValue("username") {
			delete(sessions, k)
		}
	}
	// Create a new random session token
	// we use the "github.com/google/uuid" library to generate UUIDs
	sessionToken := uuid.Must(uuid.NewV4()).String()
	expiresAt := time.Now().Add(600 * time.Second)

	// Set the token in the session map, along with the session information
	sessions[sessionToken] = session{
		username: r.FormValue("username"),
		expiry:   expiresAt,
	}

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})
	cookies[sessionToken] = getIP(r)
}

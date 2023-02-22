package session

import (
	"net/http"
	"time"
)

var sessions = map[string]session{}
var cookies = map[string]string{}

// each session contains the username of the user and the time at which it expires
type session struct {
	username string
	expiry   time.Time
}

// Checks if the session has expired
func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

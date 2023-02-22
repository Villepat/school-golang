package handlers

import (
	"encoding/json"
	"forum/server/data"
	"forum/server/database"
	"forum/server/pages"
	"net/http"
	"net/mail"
	"unicode"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pages.Register(w, r)
		return
	}
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		_, err := mail.ParseAddress(r.FormValue("email"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{"Invalid email address", 400})
			return
		}
		email := r.FormValue("email")
		if !database.IfUserExists(username, email) {
			var i data.RegisterStruct
			i.User = username
			i.Password = r.FormValue("password")
			i.Email = email
			if !isUsername(i.User) {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(&Response{"Only maximum of 10 letters allowed in username", 400})
				return
			}
			if !password(i.Password) {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(&Response{"Password must be at least 8 characters long and contain at least one number, one uppercase letter and one special character", 400})
				return
			}
			database.AddUser(i, database.Db)
			json.NewEncoder(w).Encode(&Response{"User added", 200})
			return

		}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(&Response{"Username or email already registered", 403})
		return
	}
}
func password(pass string) bool {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || !sym || tot < 8 {
		return false
	}

	return true
}

// The net/mail package implements and follows the RFC 5322 specification
// (and extension by RFC 6532). This means a seemingly bad email address
//
//	like bad-example@t is accepted and parsed by the package
//
// because it's valid according to the spec. t may be a valid local domain name,
//
//	it does not necessarily have to be a public domain. net/mail does not check
//	if the domain part of the address is a public domain,
//	nor that it is an existing, reachable public domain.

// isUsername returns true if s contains only letters and is at most 10 characters long.
func isUsername(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return len(s) <= 10
}

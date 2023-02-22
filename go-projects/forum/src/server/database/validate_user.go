package database

import (
	"database/sql"
	"log"
	"net/http"
)

type LoginStruct struct {
	User     string
	Password string
	Email    string
}

// Receives user information and checks if user credentials match to the ones stored in database.
// Id and Creation-timestamp are being queried as well, but not used, as all fields of a table have to be queried to execute a successul data fetch.
func ValidateUser(r *http.Request, db *sql.DB) bool {
	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var Username string
		var Email string
		var Password string
		var Created string
		err = rows.Scan(&id, &Username, &Email, &Password, &Created)

		if err != nil {
			log.Panic(err)
		}

		if Username == r.FormValue("username") && Password == r.FormValue("password") {
			//fmt.Println("Your credentials have been verified, try your best to contain your excitement")
			return true

		}
	}
	return false
}

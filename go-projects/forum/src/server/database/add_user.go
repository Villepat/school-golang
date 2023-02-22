package database

import (
	"database/sql"
	"forum/server/data"
	"log"
	"time"
)

func IfUserExists(user, email string) bool {
	query, err := Db.Prepare("SELECT Username, Email FROM users WHERE Username=? OR Email=?")
	if err != nil {
		log.Fatal(err)
	}

	var userInDB string
	var mailInDB string
	query.QueryRow(user, email).Scan(&userInDB, &mailInDB)
	if userInDB == user || mailInDB == email {
		return true
	} else {
		return false
	}

}

func AddUser(info data.RegisterStruct, db *sql.DB) {
	dt := time.Now().Format("2006-01-02 15:04:05")
	query, err := db.Prepare("INSERT INTO users (Username, Email, Password, Created) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	query.Exec(info.User, info.Email, info.Password, dt)
}

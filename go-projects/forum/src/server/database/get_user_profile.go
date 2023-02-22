package database

import (
	"log"
)

func GetUserProfile(user string) []string {
	result := []string{}
	getUserQuery := "SELECT Email, Created FROM users WHERE Username=?"
	rows, err := Db.Query(getUserQuery, user)

	if err != nil {
		log.Panic(err)
		return result
	}
	defer rows.Close()
	for rows.Next() {
		var email string
		var created string
		err = rows.Scan(&email, &created)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, user, email, created)

	}
	return result
}

package database

import (
	"database/sql"
	"forum/server/data"
	"log"
)

func GetCreationDates(db *sql.DB) []data.CreationDates {
	var posts []data.CreationDates
	rows, err := db.Query("SELECT id, created FROM posts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var created string
		err := rows.Scan(&id, &created)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, data.CreationDates{PostID: id, Creation: created})
	}
	return posts
}

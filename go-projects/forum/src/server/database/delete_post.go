package database

import (
	"database/sql"
	"log"
)

func DeletePost(id int, db *sql.DB) {
	_, err := db.Exec("DELETE FROM posts WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("DELETE FROM comments WHERE post_id=?", id)
	if err != nil {
		log.Fatal(err)
	}

}

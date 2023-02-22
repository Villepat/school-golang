package database

import (
	"log"
)

func DeleteComment(id int) {
	_, err := Db.Exec("DELETE FROM comments WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}
}

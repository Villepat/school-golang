package database

import (
	"database/sql"
	"log"
)

func deleteLike(post_id, comment_id int, username string) error {
	query, err := Db.Prepare("DELETE FROM likes WHERE post_id=? AND comment_id=? AND username=?")
	if err != nil {
		return err
	}
	query.Exec(post_id, comment_id, username)
	return nil
}

func addLike(status int, username string, comment_id int, post_id int, db *sql.DB) error {
	query, err := db.Prepare("INSERT INTO likes (status, username, comment_id, post_id) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
	query.Exec(status, username, comment_id, post_id)
	return nil
}

func checkIfExists(status int, username string, comment_id int, post_id int, db *sql.DB) bool {
	statement, err := db.Prepare("SELECT id FROM likes WHERE status = ? AND username = ? AND comment_id = ? AND post_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	var id int
	err = statement.QueryRow(status, username, comment_id, post_id).Scan(&id)
	if err == sql.ErrNoRows {
		return false
	}
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func AddLikeDislike(status int, username string, post_id int, comment_id int, Db *sql.DB) error {
	if checkIfExists(status, username, comment_id, post_id, Db) {
		deleteLike(post_id, comment_id, username)
		return nil
	}
	if checkIfExists(status*-1, username, comment_id, post_id, Db) {
		deleteLike(post_id, comment_id, username)
		if err := addLike(status, username, comment_id, post_id, Db); err != nil {
			return err
		}
		return nil
	}
	if err := addLike(status, username, comment_id, post_id, Db); err != nil {
		return err
	}
	return nil
}

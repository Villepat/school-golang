package database

import (
	"time"
)

func AddComment(author string, content string, responding_to_post int) error {
	dt := time.Now().Format("2006-01-02 15:04:05")
	query, err := Db.Prepare("INSERT INTO comments (author, content, created, post_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = query.Exec(author, content, dt, responding_to_post)
	if err != nil {
		return err
	}
	return nil
}

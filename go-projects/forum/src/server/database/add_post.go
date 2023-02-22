package database

import (
	"database/sql"
	"strings"
	"time"
)

func AddPost(author string, title string, content string, tags string, db *sql.DB) error {
	dt := time.Now().Format("2006-01-02 15:04:05")
	tags = strings.ReplaceAll(tags, " ", "")
	query, err := db.Prepare("INSERT INTO posts (author, title, content, created, views, tags) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = query.Exec(author, title, content, dt, 0, tags)
	if err != nil {
		return err
	}
	return nil
}

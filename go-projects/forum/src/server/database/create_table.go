package database

import (
	"database/sql"
	"log"
)

func createTables(db *sql.DB) {
	users_table := `CREATE TABLE IF NOT EXISTS users (
	    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	    "Username" TEXT,
		"Email" TEXT,
		"Password" TEXT,
		"Created" TEXT)`
	query, err := db.Prepare(users_table)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()
	post_table := `CREATE TABLE IF NOT EXISTS posts (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"author" TEXT,
		"title" TEXT,
		"content" TEXT,
		"created" TEXT,
		"views" INT NOT NULL, 
		"tags" TEXT 
	)`
	query, err = db.Prepare(post_table)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()
	comment_table := `CREATE TABLE IF NOT EXISTS comments (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"author" TEXT,
		"content" TEXT,
		"created" TEXT,
		"post_id" INTEGER NOT NULL,
		CONSTRAINT fk_posts
			FOREIGN KEY (post_id)
			REFERENCES posts (id)
			ON DELETE CASCADE
	)`
	query, err = db.Prepare(comment_table)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()
	likes_table := `CREATE TABLE IF NOT EXISTS likes (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"status" INTEGER,
		"username" TEXT,
		"comment_id" INTEGER,
		"post_id" INTEGER
	)`
	query, err = db.Prepare(likes_table)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()
}

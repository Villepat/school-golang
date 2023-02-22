package database

import (
	"database/sql"
	"log"
	"strconv"
)

func IncrementViews(post_id int, db *sql.DB) {
	row, err := db.Query("SELECT * FROM posts WHERE id=?", post_id)
	if err != nil {
		log.Panic(err)
	}
	var id int
	var author string
	var title string
	var content string
	var created string
	var views int
	var tags string
	defer row.Close()
	for row.Next() {
		err = row.Scan(&id, &author, &title, &content, &created, &views, &tags)
		if err != nil {
			log.Panic(err.Error())
		}
	}
	command := "UPDATE posts SET views=" + strconv.Itoa(views+1) + " WHERE id=" + strconv.Itoa(post_id)
	query, err := db.Prepare(command)
	if err != nil {
		log.Panic(err.Error())
	}
	query.Exec()
}

package database

import (
	"forum/server/data"
	"log"
)

func GetComments(post_id int) []data.Comment {
	result := []data.Comment{}

	rows, err := Db.Query("SELECT * FROM comments WHERE post_id=? ORDER BY created DESC", post_id)
	if err != nil {
		log.Fatal(err)
		return result
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var Author string
		var Content string
		var Created string
		var PostId int
		err := rows.Scan(&id, &Author, &Content, &Created, &PostId)
		if err != nil {
			log.Fatal(err)
		}
		comment := data.Comment{
			Id:           id,
			Author:       Author,
			Content:      Content,
			CreationDate: Created,
		}
		result = append(result, comment)
	}
	return result
}

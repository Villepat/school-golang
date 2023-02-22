package database

import (
	"forum/server/data"
	"log"
)

func getCommentIds(user string) []int {
	statement, err := Db.Prepare("SELECT comment_id FROM likes WHERE username=? AND status=1")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := statement.Query(user)
	if err != nil {
		log.Fatal(err)
	}
	comment_ids := []int{}
	defer rows.Close()
	for rows.Next() {
		var comment_id int
		err := rows.Scan(&comment_id)
		if err != nil {
			log.Fatal(err)
		}
		comment_ids = append(comment_ids, comment_id)
	}
	return comment_ids
}

func getPostIds(comment_ids []int) []int {
	ids := []int{}
	for _, v := range comment_ids {
		statement, err := Db.Prepare("SELECT post_id FROM comments WHERE id=?")
		if err != nil {
			log.Fatal(err)
		}
		rows, err := statement.Query(v)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			err := rows.Scan(&id)
			if err != nil {
				log.Fatal(err)
			}
			ids = append(ids, id)
		}
	}
	return ids
}

func existsInArr(arr []int, i int) bool {
	for _, v := range arr {
		if v == i {
			return true
		}
	}
	return false
}

func removeDuplicates(arr []int) []int {
	result := []int{}
	for _, v := range arr {
		if !existsInArr(result, v) {
			result = append(result, v)
		}
	}
	return result
}

func GetPostThroughLikedComments(user string) ([]data.Post, error) {
	ids := getPostIds(getCommentIds(user))
	ids = removeDuplicates(ids)
	result := []data.Post{}
	var err error
	for _, v := range ids {
		statement, err := Db.Prepare("SELECT id, title, author FROM posts WHERE id=?")
		if err != nil {
			log.Println(err)
			return result, err
		}
		rows, err := statement.Query(v)
		if err != nil {
			log.Println(err)
			return result, err
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var title string
			var author string
			err := rows.Scan(&id, &title, &author)
			url := MakeURL(title, "")
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, data.Post{Id: id, Title: title, Author: author, URL: url})
		}
	}
	return result, err
}

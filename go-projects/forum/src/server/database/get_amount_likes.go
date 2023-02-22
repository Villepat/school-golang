package database

import "forum/server/data"

func GetAmountOfLikes(post data.Post) (data.Post, error) {
	statement, err := Db.Prepare("SELECT * FROM likes WHERE post_id=?")
	if err != nil {
		return data.Post{}, err
	}
	rows, err := statement.Query(post.Id)
	if err != nil {
		return data.Post{}, err
	}
	var id int
	var status int
	var username string
	var comment_id int
	var post_id int
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &status, &username, &comment_id, &post_id)
		if err != nil {
			return data.Post{}, err
		}
		if status == 1 {
			post.Likes++
			continue
		}
		post.Dislikes++
	}
	for i, v := range post.Comments {
		statement, err := Db.Prepare("SELECT * FROM likes WHERE comment_id=?")
		if err != nil {
			return data.Post{}, err
		}
		rows, err := statement.Query(v.Id)
		if err != nil {
			return data.Post{}, err
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &status, &username, &comment_id, &post_id)
			if err != nil {
				return data.Post{}, err
			}
			if status == 1 {
				post.Comments[i].Likes++
				continue
			}
			post.Comments[i].Dislikes++
		}
	}
	return post, nil
}

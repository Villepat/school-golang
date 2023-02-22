package database

import (
	"database/sql"
	"forum/server/data"
)

func getAllLikesByName(username string, db *sql.DB) ([]data.Like, error) {
	statement, err := db.Prepare("SELECT * FROM likes WHERE username=?")
	if err != nil {
		return []data.Like{}, err
	}
	rows, err := statement.Query(username)
	if err != nil {
		return []data.Like{}, err
	}
	res := []data.Like{}
	defer rows.Close()
	for rows.Next() {
		res = append(res, data.Like{})
		i := len(res) - 1
		err := rows.Scan(&res[i].Id, &res[i].Status, &res[i].Username, &res[i].CommentId, &res[i].PostId)
		if err != nil {
			return res, err
		}
	}
	return res, nil
}

func getSpecificLike(comment_id int, arr []data.Like) data.Like {
	for _, v := range arr {
		if v.CommentId == comment_id {
			return v
		}
	}
	return data.Like{}
}

func GetLikes(username string, db *sql.DB, post data.Post) (data.Post, error) {
	all_likes, err := getAllLikesByName(username, db)
	if err != nil {
		return data.Post{}, err
	}
	for _, v := range all_likes {
		if v.PostId == post.Id {
			post.LikeStatus = v.Status
		}
	}
	for i, v := range post.Comments {
		like := getSpecificLike(v.Id, all_likes)
		if like.Id == 0 {
			continue
		}
		post.Comments[i].LikeStatus = like.Status
	}
	return post, nil
}

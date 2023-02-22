package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

var Rows *sql.Rows

func GetTags(id int) ([]string, error) {
	var tags []string
	var err error

	if id == 0 {

		Rows, err = Db.Query("SELECT DISTINCT tags FROM posts")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		tags = tagscan()

	} else {
		Rows, err = Db.Query("SELECT DISTINCT tags FROM posts WHERE id = ?", id)
		if err != nil {
			log.Panic(err)
		}
		tags = tagscan()

	}

	// remove empty strings
	list := []string{}
	for _, entry := range tags {
		if entry != "" {
			list = append(list, entry)
		}
	}

	return list, nil
}

func tagscan() []string {
	tags := []string{}
	defer Rows.Close()
	for Rows.Next() {
		var tag string
		err := Rows.Scan(&tag)
		if err != nil {
			log.Panic(err)
		}
		tag = strings.ReplaceAll(tag, " ", "")
		tagslice := strings.Split(tag, ",")

		tags = append(tags, tagslice...)

	}
	return tags
}

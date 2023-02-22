package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func init() {
	OpenDatabase()
	createTables(Db)
}

func OpenDatabase() {
	dbLocal, err := sql.Open("sqlite3", "foo.db")
	if err != nil {
		log.Panic(err)
	}
	Db = dbLocal
}

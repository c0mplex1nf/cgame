package deck

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/mattn/go-sqlite3"
)

func (db *Db) connect() db {
	db, err := sql.Open("sqlite3", "./deck.db")
	checkErr(err)
	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
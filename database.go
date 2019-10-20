package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var database Database

type Database struct {
	Conn *sql.DB
}

func (db *Database) Init() {
	var err error
	db.Conn, err = sql.Open("sqlite3", config.DatabaseFileName)
	if err != nil {
		panic(err)
	}
	_, err = db.Conn.Exec(`
		create table if not exists log (
            id integer primary key autoincrement,
            time datetime,
            type text,
            value text
		);
	`)
	if err != nil {
		panic(err)
	}
}

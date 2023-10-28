package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// CREATE TABLE <table_name>(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT
// NULL, name INTEGER DEFAULT 0);
// INSERT INTO <table_name>(<column_name>, <column_name>) VALUES('text', int);
// DROP TABLE <table_name>

const db_name = "./database.sqlite3"

func main() {

	sqlStmt := `create table if not exists <table_name>(<column_name>, <column_name>) values(<value>, <value>)`
	db, err := sql.Open("sqlite3", db_name)

	if err != nil {
		log.Printf("%q:\n", err)
	}

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q:\n", err)
	}

	defer db.Close()
}

// sqlStmt := fmt.Sprintf("statement %s, %d", variable, variable)

package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type DB struct {
	Conn *sql.DB
}

func (d *DB) Init(dbPath string) {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		f, err := os.Create(dbPath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	d.Conn = db
	d.CreateTables()
}

func (d *DB) CreateTables() {
	// registered users table
	stmt := `CREATE TABLE if not exists users (
		"uid" TEXT primary key,
		"log" TEXT,
		"pas" TEXT
	);`
	_, err := d.Conn.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}

	// tickets table
	stmt = `CREATE TABLE if not exists tickets (
		"uid" TEXT primary key,
		"from"	TEXT,
		"to" TEXT,
		"date" TEXT,
		"departure" TEXT,
		"arrival" TEXT,
		"transfer" TEXT,
		"tr_departure" TEXT,
		"time" TEXT
	);`
	_, err = d.Conn.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}

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
	// peoples point table
	stmt := `CREATE TABLE if not exists points (
		"uid" TEXT primary key,
		"name" TEXT,
		"art" TEXT,
		"points" NUMBER 
	);`
	_, err := d.Conn.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}

	// famous art table
	stmt = `CREATE TABLE if not exists arts (
		"uid" TEXT primary key,
		"city" TEXT,
		"art" TEXT,
-- 		"opentime" TEXT,
-- 		"closetime" TEXT,
		"intime" TEXT,
		"tag" TEXT
	);`
	_, err = d.Conn.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}

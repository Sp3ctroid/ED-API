package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func OpenStorage() *sql.DB {

	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "albums",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	return db
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to DB!")
}

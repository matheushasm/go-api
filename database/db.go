package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error

	dsn := "root:password@tcp(127.0.0.1:3306)/my_database"

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error during database connection:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Error during database verification:", err)
	}

	log.Println("Connected to database!")
}

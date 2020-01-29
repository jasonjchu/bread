package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var Database *sql.DB

func OpenConnection() error {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Panicf("Error: Failed to establish database connection %v", err)
	}
	Database = db
	return nil
}

func CloseConnection() {
	Database.Close()
}
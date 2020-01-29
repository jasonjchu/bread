package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/jasonjchu/bread/app/env"
	"log"
	"os"
)

var Pool *sql.DB

func OpenConnection() error {
	dsn := mysql.NewConfig()
	dsn.User = os.Getenv(env.DBUserKey)
	dsn.Passwd = os.Getenv(env.DBPasswordKey)
	dsn.Net = os.Getenv(env.DBNetKey) // Connection type, required for [dsn.Addr]
	dsn.Addr = os.Getenv(env.DBHostKey)
	dsn.DBName = os.Getenv(env.DBNameKey)

	db, err := sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		log.Panicf("Error: Failed to establish database connection %v", err)
	}
	Pool = db
	return nil
}

func CloseConnection() {
	Pool.Close()
}

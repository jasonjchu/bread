package db

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jasonjchu/bread/app/env"
	"github.com/jmoiron/sqlx"
	"os"
)

var Pool *sqlx.DB

func OpenConnection() error {
	dsn := mysql.NewConfig()
	dsn.User = os.Getenv(env.DBUserKey)
	dsn.Passwd = os.Getenv(env.DBPasswordKey)
	dsn.Net = os.Getenv(env.DBNetKey) // Connection type, required for [dsn.Addr]
	dsn.Addr = os.Getenv(env.DBHostKey)
	dsn.DBName = os.Getenv(env.DBNameKey)

	// Make sure we can parse DATE into time.Time
	dsn.Params = make(map[string]string)
	dsn.Params["parseTime"] = "true"

	pool, err := sqlx.Connect("mysql", dsn.FormatDSN())
	if err != nil {
		return err
	}
	Pool = pool
	return nil
}

func CloseConnection() {
	Pool.Close()
}

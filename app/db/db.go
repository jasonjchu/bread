package db

import (
	"encoding/json"
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

	// Preprocess AWS creds from secret if they exist
	_, credsExist := os.LookupEnv(env.DBCreds)
	if credsExist {
		dbCreds := make(map[string]string)
		json.Unmarshal([]byte(os.Getenv(env.DBCreds)), &dbCreds)
		dsn.Passwd = dbCreds["password"]
	}

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

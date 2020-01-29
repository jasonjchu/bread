package env

import (
	"github.com/joho/godotenv"
)

const (
	DBHostKey     string = "BREAD_DB_HOST"
	DBUserKey     string = "BREAD_DB_USER"
	DBPasswordKey string = "BREAD_DB_PASSWD"
	DBNameKey     string = "BREAD_DB_NAME"
	DBNetKey      string = "BREAD_DB_NET"
)

func LoadEnv() error {
	return godotenv.Load()
}

package account

import "github.com/jasonjchu/bread/app/db"

type Id int64

func CreateAccount(username string, password string) (Id, error) {
	pool := db.Pool
	insertQuery := "INSERT INTO accounts (username, password) VALUES (?, ?)"
	res, err := pool.Exec(insertQuery, username, password)
	if err != nil {
		return -1, err
	}
	accountId, err := res.LastInsertId()
	return Id(accountId), err
}


package account

import (
	"errors"
	"github.com/jasonjchu/bread/app/db"
	"github.com/jmoiron/sqlx"
)

type Id int64
type Username string
type Password string

type Account struct {
	Id       Id       `db:"_id"`
	Username Username `db:"username"`
	Password Password `db:"password"`
}

func CreateAccount(username Username, password Password) (Id, error) {
	pool := db.Pool
	insertQuery := "INSERT INTO accounts (username, password) VALUES (?, ?)"
	res, err := pool.Exec(insertQuery, username, password)
	if err != nil {
		return -1, err
	}
	accountId, err := res.LastInsertId()
	return Id(accountId), err
}

func GetAccountByUsername(username Username) (*Account, error) {
	pool := db.Pool
	row := pool.QueryRowx("SELECT * FROM accounts WHERE username=?", username)

	account, err := scanAccountFromRow(row)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func VerifyPassword(accountPassword Password, password Password) error {
	if password != accountPassword {
		return errors.New("invalid password for account")
	}
	return nil
}

func scanAccountFromRow(row *sqlx.Row) (*Account, error) {
	account := Account{}
	err := row.StructScan(&account)
	if err != nil {
		return nil, err
	}
	return &account, err
}

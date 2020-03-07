package candidate

import (
	"github.com/jasonjchu/bread/app/db"
	"github.com/jasonjchu/bread/app/models/account"
	"time"
)

func CreateCandidate(
	name string,
	program string,
	gradDate time.Time,
	description string,
	accountId account.Id,
) error {
	pool := db.Pool
	insertQuery := "INSERT INTO candidates (_id, name, program, grad_date, description) VALUES (?, ?, ?, ?, ?)"
	res, err := pool.Exec(insertQuery, accountId, name, program, gradDate, description)
	if err != nil {
		return err
	}
	_, err = res.LastInsertId()
	return err
}

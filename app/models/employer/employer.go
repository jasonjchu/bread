package employer

import (
	"github.com/jasonjchu/bread/app/db"
	"github.com/jasonjchu/bread/app/models/account"
	"github.com/jmoiron/sqlx"
)

type Id int
type Name string
type WorksAt int

type Employer struct {
	Id      Id      `db:"_id"`
	Name    Name    `db:"name"`
	WorksAt WorksAt `db:"works_at"`
}

func GetEmployerById(id Id) (*Employer, error) {
	pool := db.Pool
	row := pool.QueryRowx("SELECT * FROM employers WHERE _id=?", id)

	employer, err := scanEmployerFromRow(row)
	if err != nil {
		return nil, err
	}
	return employer, nil
}

func scanEmployerFromRow(row *sqlx.Row) (*Employer, error) {
	employer := Employer{}
	err := row.StructScan(&employer)
	if err != nil {
		return nil, err
	}
	return &employer, err
}

func CreateEmployer(name Name, worksAt WorksAt, accountId account.Id) error {
	pool := db.Pool
	insertQuery := "INSERT INTO employers (_id, name, works_at) VALUES (?, ?, ?)"
	res, err := pool.Exec(insertQuery, accountId, name, worksAt)
	if err != nil {
		return err
	}
	_, err = res.LastInsertId()
	return err
}

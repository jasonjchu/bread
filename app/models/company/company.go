package company

import (
	"github.com/jasonjchu/bread/app/db"
	"github.com/jmoiron/sqlx"
)

type Id int
type Name string
type Description string

type Companies []*Company

type Company struct {
	Id          Id          `db:"_id"`
	Name        Name        `db:"name"`
	Description Description `db:"description"`
}

func GetCompaniesWithName(name string) (Companies, error) {
	var selectQuery string
	var rows *sqlx.Rows
	var err error
	pool := db.Pool
	if name == "" {
		selectQuery = "SELECT * FROM companies"
		rows, err = pool.Queryx(selectQuery)
	} else {
		selectQuery = "SELECT * FROM companies WHERE UPPER(name) LIKE UPPER(CONCAT('%', ?, '%'))"
		rows, err = pool.Queryx(selectQuery, name)
	}

	if err != nil {
		return nil, err
	}

	companies, err := scanCompaniesFromRows(rows)
	if err != nil {
		return nil, err
	}
	return companies, nil
}

func scanCompanyFromRow(row *sqlx.Row) (*Company, error) {
	company := Company{}
	err := row.StructScan(&company)
	if err != nil {
		return nil, err
	}
	return &company, err
}

func scanCompaniesFromRows(rows *sqlx.Rows) (Companies, error) {
	var companies Companies
	var err error
	for rows.Next() {
		company := Company{}
		err = rows.StructScan(&company)
		if err != nil {
			return nil, err
		}
		companies = append(companies, &company)
	}
	return companies, err
}

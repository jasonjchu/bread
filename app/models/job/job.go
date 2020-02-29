package job

import (
	"database/sql"
	"github.com/jasonjchu/bread/app/db"
	"github.com/jmoiron/sqlx"
)

type Id string
type CompanyId int
type Country string
type CountryCode string
type Board string
type Description string
type Title string
type Type string
type Location string
type Organization string
type URL string
type Salary string
type Sector string

type Jobs []*Job

// TODO(kallentu): Make component (usable) struct for Job.
type Job struct {
	Id           Id           `db:"_id"`
	CompanyId    CompanyId    `db:"company_id"`
	Country      Country      `db:"country"`
	CountryCode  CountryCode  `db:"country_code"`
	DateAdded    sql.NullTime `db:"date_added"` // Time is optional, must check if [DateAdded.Valid] before using
	HasExpired   bool         `db:"has_expired"`
	Board        Board        `db:"job_board"`
	Description  Description  `db:"job_description"`
	Title        Title        `db:"job_title"`
	Type         Type         `db:"job_type"`
	Location     Location     `db:"location"`
	Organization Organization `db:"organization"`
	URL          URL          `db:"page_url"`
	Salary       Salary       `db:"salary"`
	Sector       Sector       `db:"sector"`
}

func GetJobById(id Id) (*Job, error) {
	pool := db.Pool
	row := pool.QueryRowx("SELECT * FROM jobs WHERE _id=?", id)

	job, err := scanJobFromRow(row)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func GetJobs(numberOfJobs int) (Jobs, error) {
	pool := db.Pool
	rows, err := pool.Queryx("SELECT * FROM jobs LIMIT ?", numberOfJobs)
	if err != nil {
		return nil, err
	}

	jobs, err := scanJobsFromRows(rows)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func scanJobFromRow(row *sqlx.Row) (*Job, error) {
	job := Job{}
	err := row.StructScan(&job)
	if err != nil {
		return nil, err
	}
	return &job, err
}

// Similar logic to [scanJobFromRow], but using different struct [sql.Rows].
func scanJobsFromRows(rows *sqlx.Rows) (Jobs, error) {
	var jobs Jobs
	var err error
	for rows.Next() {
		job := Job{}
		err = rows.StructScan(&job)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, &job)
	}
	return jobs, err
}


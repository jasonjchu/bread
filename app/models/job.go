package models

import (
	"database/sql"
	"github.com/jasonjchu/bread/app/db"
	"time"
)

type Id string
type Country string
type CountryCode string
type DateAdded time.Time
type HasExpired bool
type Board string
type Description string
type Title string
type Type string
type Location string
type Organization string
type URL string
type Salary string
type Sector string

type Job struct {
	Id           Id           `json:"uniq_id"`
	Country      Country      `json:"country"`
	CountryCode  CountryCode  `json:"country_code"`
	DateAdded    DateAdded    `json:"date_added"`
	HasExpired   HasExpired   `json:"has_expired"`
	Board        Board        `json:"job_board"`
	Description  Description  `json:"job_description"`
	Title        Title        `json:"job_title"`
	Type         Type         `json:"job_type"`
	Location     Location     `json:"location"`
	Organization Organization `json:"organization"`
	URL          URL          `json:"page_url"`
	Salary       Salary       `json:"salary"`
	Sector       Sector       `json:"sector"`
}

func GetJobById(id Id) (*Job, error) {
	pool := db.Pool
	row := pool.QueryRow("SELECT * FROM jobs WHERE uniq_id=?", id)

	job, err := scanJobFromRow(row)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func GetJobs(numberOfJobs int) ([]*Job, error) {
	pool := db.Pool
	rows, err := pool.Query("SELECT * FROM jobs LIMIT ?", numberOfJobs)

	jobs, err := scanJobFromRows(rows)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func scanJobFromRow(row *sql.Row) (*Job, error) {
	var job Job
	var dateAdded sql.NullTime
	var hasExpired int64
	err := row.Scan(&job.Id,
		&job.Country,
		&job.CountryCode,
		&dateAdded,
		&hasExpired,
		&job.Board,
		&job.Description,
		&job.Title,
		&job.Type,
		&job.Location,
		&job.Organization,
		&job.URL,
		&job.Salary,
		&job.Sector)
	if err != nil {
		return nil, err
	}
	fillExtraFields(&job, dateAdded, hasExpired)
	return &job, err
}

// Similar logic to [scanJobFromRow], but using different struct [sql.Rows].
// Unfortunately, code duplication is necessary even though behaviour is similar.
func scanJobFromRows(rows *sql.Rows) ([]*Job, error) {
	var jobs []*Job
	var err error
	for rows.Next() {
		var job Job
		var dateAdded sql.NullTime
		var hasExpired int64
		err = rows.Scan(&job.Id,
			&job.Country,
			&job.CountryCode,
			&dateAdded,
			&hasExpired,
			&job.Board,
			&job.Description,
			&job.Title,
			&job.Type,
			&job.Location,
			&job.Organization,
			&job.URL,
			&job.Salary,
			&job.Sector)
		if err != nil {
			return nil, err
		}
		fillExtraFields(&job, dateAdded, hasExpired)
		jobs = append(jobs, &job)
	}

	return jobs, err
}

// Jobs need additional processing for certain fields when converting from MySQL data types.
func fillExtraFields(job *Job, dateAdded sql.NullTime, hasExpired int64) {
	// [date_added] is an optional field.
	// If we find a non-null time, populate field in Job.
	if dateAdded.Valid {
		job.DateAdded = DateAdded(dateAdded.Time)
	}

	// [has_expired] will be an int64 value when scanned from the database; we need to do a quick conversion.
	if hasExpired == 0 {
		job.HasExpired = false
	} else {
		job.HasExpired = true
	}
}

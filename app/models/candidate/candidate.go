package candidate

import (
	"github.com/jasonjchu/bread/app/db"
	"github.com/jasonjchu/bread/app/models/account"
	"github.com/jmoiron/sqlx"
	"time"
)

type Id int
type Name string
type Program string
type Description string
type Candidates []*Candidate

type Candidate struct {
	Id          Id          `db:"_id"`
	Name        Name        `db:"name"`
	Program     Program     `db:"program"`
	GradDate    time.Time   `db:"grad_date"`
	Description Description `db:"description"`
}

func GetCandidateById(id Id) (*Candidate, error) {
	pool := db.Pool
	row := pool.QueryRowx("SELECT * FROM candidates WHERE _id=?", id)

	candidate, err := scanCandidateFromRow(row)
	if err != nil {
		return nil, err
	}
	return candidate, nil
}

func GetCandidatesByJidLiked(id string, candidateLimit int) (Candidates, error){
	pool := db.Pool
	rows, err := pool.Queryx("SELECT * FROM candidates WHERE _id in " +
		"(SELECT cid FROM candidateSeenJob WHERE jid=? AND liked=True AND cid NOT IN " +
		"(SELECT cid FROM jobSeenCandidate WHERE jid=?)) LIMIT ?", id, id, candidateLimit)
	if err != nil {
		return nil, err
	}
	candidates, err := scanCandidateFromRows(rows)
	if err != nil {
		return nil, err
	}
	return candidates, nil
}

func scanCandidateFromRow(row *sqlx.Row) (*Candidate, error) {
	candidate := Candidate{}
	err := row.StructScan(&candidate)
	if err != nil {
		return nil, err
	}
	return &candidate, err
}

func scanCandidateFromRows(rows *sqlx.Rows) (Candidates, error) {
	var candidates Candidates
	var err error
	for rows.Next() {
		candidate := Candidate{}
		err = rows.StructScan(&candidate)
		if err != nil {
			return nil, err
		}
		candidates = append(candidates, &candidate)
	}
	return candidates, err
}

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

package candidateSeenJob

import (
	"github.com/jasonjchu/bread/app/db"
	"github.com/jasonjchu/bread/app/models/candidate"
	"github.com/jasonjchu/bread/app/models/job"
)

type CandidateSeenJob struct {
	CandidateId candidate.Id `db:"cid"`
	JobId       job.Id       `db:"jid"`
	Liked       bool         `dbb:"liked"`
}

func insertCandidateSeenJob(cid candidate.Id, jid job.Id, liked bool) error {
	pool := db.Pool
	insertQuery := `INSERT INTO candidateSeenJob (cid, jid, liked) VALUES (?, ?, ?)`
	_, err := pool.Exec(insertQuery, cid, jid, liked)
	return err
}

func InsertCandidateLikesJob(cid candidate.Id, jid job.Id) error {
	return insertCandidateSeenJob(cid, jid, true)
}

func InsertCandidateDislikesJob(cid candidate.Id, jid job.Id) error {
	return insertCandidateSeenJob(cid, jid, false)
}

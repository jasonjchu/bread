package jobSeenCandidate

import (
	"github.com/jasonjchu/bread/app/db"
	"github.com/jasonjchu/bread/app/models/candidate"
	"github.com/jasonjchu/bread/app/models/job"
	"github.com/jasonjchu/bread/app/models/match"
)

type JobSeenCandidate struct {
	JobId       job.Id       `db:"jid"`
	CandidateId candidate.Id `db:"cid"`
	Liked       bool         `db:"liked"`
}

func insertJobSeenCandidate(jid job.Id, cid candidate.Id, liked bool) error {
	pool := db.Pool
	insertQuery := `INSERT INTO jobSeenCandidate (cid, jid, liked) VALUES (?, ?, ?)`
	_, err := pool.Exec(insertQuery, cid, jid, liked)
	return err
}

func InsertJobLikesCandidate(jid job.Id, cid candidate.Id) error {
	err := insertJobSeenCandidate(jid, cid, true)
	if err != nil {
		return err
	}

	// Under the assumption that in order for a Job to like a Candidate means that the Candidate must have already liked
	// the Job, this creates a new Match.
	err = match.InsertMatch(cid, jid)
	return err
}

func InsertJobDislikesCandidate(jid job.Id, cid candidate.Id) error {
	return insertJobSeenCandidate(jid, cid, false)
}

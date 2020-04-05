package employerDislikesCandidateService

import (
	"github.com/jasonjchu/bread/app/models/candidate"
	"github.com/jasonjchu/bread/app/models/job"
	"github.com/jasonjchu/bread/app/models/jobSeenCandidate"
)

type Request struct {
	JobId       job.Id       `json:"job_id"`
	CandidateId candidate.Id `json:"candidate_id"`
}

func Exec(req Request) error {
	return jobSeenCandidate.InsertJobDislikesCandidate(req.JobId, req.CandidateId)
}

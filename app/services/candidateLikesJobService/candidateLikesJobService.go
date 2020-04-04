package candidateLikesJobService

import (
	"github.com/jasonjchu/bread/app/models/candidate"
	"github.com/jasonjchu/bread/app/models/candidateSeenJob"
	"github.com/jasonjchu/bread/app/models/job"
)

type Request struct {
	CandidateId candidate.Id `json:"candidate_id"`
	JobId       job.Id       `json:"job_id"`
}

func Exec(req Request) error {
	return candidateSeenJob.InsertCandidateLikesJob(req.CandidateId, req.JobId)
}

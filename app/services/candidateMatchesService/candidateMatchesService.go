package candidateMatchesService

import (
	"github.com/jasonjchu/bread/app/models/candidate"
	"github.com/jasonjchu/bread/app/models/company"
	"github.com/jasonjchu/bread/app/models/job"
	"github.com/jasonjchu/bread/app/models/match"
)

type Request struct {
	Id candidate.Id `json:"employer_id"`
}

type Response struct {
	Matches *Matches `json:"matches"`
}

type Matches []*JobMatch

type JobMatch struct {
	JobId       job.Id        `json:"job_id"`
	CompanyId   job.CompanyId `json:"company_id"`
	CompanyName company.Name  `json:"company_name"`
	JobTitle    job.Title     `json:"job_title"`
	JobLocation job.Location  `json:"job_location"`
}

func Exec(request Request) (*Response, error) {
	matchesData, err := match.GetMatchForCandidate(request.Id)
	if err != nil {
		return nil, err
	}

	var matches Matches
	for _, val := range matchesData {
		jobMatch := JobMatch{
			JobId:       val.JobId,
			CompanyId:   val.CompanyId,
			CompanyName: val.CompanyName,
			JobTitle:    val.JobTitle,
			JobLocation: val.JobLocation,
		}
		matches = append(matches, &jobMatch)
	}

	response := Response{Matches: &matches}
	return &response, nil
}

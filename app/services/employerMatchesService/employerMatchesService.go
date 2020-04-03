package employerMatchesService

import (
	"github.com/jasonjchu/bread/app/models/candidate"
	"github.com/jasonjchu/bread/app/models/company"
	"github.com/jasonjchu/bread/app/models/employer"
	"github.com/jasonjchu/bread/app/models/job"
	"github.com/jasonjchu/bread/app/models/match"
)


type Request struct {
	Id employer.Id `json:"employer_id"`
}

type Response struct {
	Matches *Matches `json:"matches"`
}

type Matches []*JobMatch
type Profiles []*CandidateProfile

type JobMatch struct {
	JobId       job.Id              `json:"job_id"`
	CompanyId   job.CompanyId       `json:"company_id"`
	CompanyName company.Name        `json:"company_name"`
	JobTitle    job.Title           `json:"job_title"`
	JobLocation job.Location        `json:"job_location"`
	Candidates  Profiles            `json:"candidate_profiles"`
}

type CandidateProfile struct {
	CandidateId       candidate.Id       `json:"id"`
	CandidateName     candidate.Name     `json:"name"`
	CandidateProgram  candidate.Program  `json:"program"`
	CandidateGradDate candidate.GradDate `json:"grad_date"`
}

func Exec(request Request) (*Response, error) {
	matchesData, err:= match.GetMatchForEmployer(request.Id)
	if err != nil {
		return nil, err
	}
	// transform data: group candidates by job
	jobMatchGroups := make(map[job.Id]*JobMatch)

	for _, jobMatch := range matchesData {

		candidateProfile := CandidateProfile{
			CandidateId:       jobMatch.CandidateId,
			CandidateName:     jobMatch.CandidateName,
			CandidateProgram:  jobMatch.CandidateProgram,
			CandidateGradDate: jobMatch.CandidateGradDate,
		}

		// if seen job before, just add the candidate to the profiles list
		if val, ok := jobMatchGroups[jobMatch.JobId]; ok {

			val.Candidates = append(val.Candidates, &candidateProfile)
		} else {
			// make new entry
			newEntry := JobMatch{
				JobId:       jobMatch.JobId,
				CompanyId:   jobMatch.CompanyId,
				CompanyName: jobMatch.CompanyName,
				JobTitle:    jobMatch.JobTitle,
				JobLocation: jobMatch.JobLocation,
				Candidates:  Profiles{&candidateProfile},
			}

			jobMatchGroups[jobMatch.JobId] = &newEntry
		}
	}

	// make an array of just the values in the map
	matches := Matches{}
	for _, jobMatch := range jobMatchGroups {
		matches = append(matches, jobMatch)
	}
	response := Response{Matches: &matches}
	return &response, nil
}

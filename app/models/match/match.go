package match

import (
	"github.com/jasonjchu/bread/app/db"
	"github.com/jasonjchu/bread/app/models/candidate"
	"github.com/jasonjchu/bread/app/models/company"
	"github.com/jasonjchu/bread/app/models/employer"
	"github.com/jasonjchu/bread/app/models/job"
	"github.com/jmoiron/sqlx"
)


type EmployerMatches []*EmployerMatch

// Used by Employers to see candidates who matched with their jobs
type EmployerMatch struct {
	CandidateId       candidate.Id       `db:"can_id"`
	JobId             job.Id             `db:"jid"`
	// include some basic job & company info
	CompanyId         job.CompanyId      `db:"comp_id"`
	CompanyName       company.Name       `db:"comp_name"`
	JobTitle          job.Title          `db:"job_title"`
	JobLocation       job.Location       `db:"location"`
	// include some basic candidate profile info
	CandidateName     candidate.Name     `db:"cand_name"`
	CandidateProgram  candidate.Program  `db:"program"`
	CandidateGradDate candidate.GradDate `db:"grad_date"`
}


// Same as employer match, minus candidate profile info
// Used by Candidates to see jobs they matched to
type CandidateMatch struct {
	CandidateId       candidate.Id       `db:"can_id"`
	JobId             job.Id             `db:"jid"`
	// include some basic job & company info
	CompanyId         job.CompanyId      `db:"comp_id"`
	CompanyName       company.Name       `db:"comp_name"`
	JobTitle          job.Title          `db:"job_title"`
	JobLocation       job.Location       `db:"location"`
}


func GetMatchForEmployer(empId employer.Id) (EmployerMatches, error) {
	const employerMatchQuery =
		`SELECT 
           candidates._id  as can_id,
		   jobs._id        as jid,
		   companies._id   as comp_id,
		   companies.name  as comp_name,
		   jobs.job_title,
		   jobs.location,
		   candidates.name as cand_name,
		   candidates.program,
		   candidates.grad_date
		FROM employers,
     		jobs,
     		matches,
     		candidates,
     		companies
		WHERE employers._id = ?
			AND jobs.company_id = employers.works_at
			AND jobs.company_id = companies._id
			AND matches.jid = jobs._id
			AND candidates._id = matches.uid;`

	pool := db.Pool
	rows, err := pool.Queryx(employerMatchQuery, empId)
	if err != nil {
		return nil, err
	}

	matches, err := scanEmployerMatchesFromRows(rows)
	if err != nil {
		return nil, err
	}

	return matches, err
}

func scanEmployerMatchesFromRows(rows *sqlx.Rows) (EmployerMatches, error) {
	var matches EmployerMatches
	var err error
	for rows.Next() {
		match := EmployerMatch{}
		err = rows.StructScan(&match)
		if err != nil {
			return nil, err
		}

		matches = append(matches, &match)
	}
	return matches, err
}

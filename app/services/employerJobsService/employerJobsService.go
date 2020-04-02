package employerJobsService

import (
	"github.com/jasonjchu/bread/app/models/employer"
	"github.com/jasonjchu/bread/app/models/job"
)

type Request struct {
	Id employer.Id `json:"employer_id"`
}

type Response struct {
	Jobs job.Jobs `json:"jobs"`
}

func Exec(req Request) (*Response, error) {
	emp, err := employer.GetEmployerById(req.Id)
	if err != nil {
		return nil, err
	}
	jobs, err := job.GetJobsByCompany(job.CompanyId(int(emp.WorksAt)))
	res := Response{Jobs: jobs}
	return &res, nil
}

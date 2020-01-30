package getJobsHandler

import (
	"github.com/jasonjchu/bread/app/models/job"
	"github.com/jasonjchu/bread/app/utils"
	"net/http"
)
const RouteURL string = "/"

func Handler(w http.ResponseWriter, r *http.Request) {
	// TODO: Get the context from r to specify the number of jobs / country code / has expired / etc
	// default the number of jobs to 200
	numberOfJobs := 200
	jobs, err := job.GetJobs(numberOfJobs)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write([]byte(utils.ToJson(jobs)))
}

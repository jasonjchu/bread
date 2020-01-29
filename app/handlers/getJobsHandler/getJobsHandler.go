package getJobsHandler

import (
	"github.com/jasonjchu/bread/app/models"
	"log"
	"net/http"
)
const RouteURL string = "/"

func Handler(w http.ResponseWriter, r *http.Request) {
	// TODO: Get the context from r to specify the number of jobs / country code / has expired / etc
	// default the number of jobs to 200
	numberOfJobs := 200
	jobs, err := models.GetJobs(numberOfJobs)
	if err != nil {
		// Write to http error and print to stderr.
		http.Error(w, err.Error(), 400)
		log.Printf("Error: Failed to load environment variables %v", err)
		return
	}
	// move this logic to the job table
	jobsToString, err := models.StringifyJobs(jobs)
	if err != nil {
		// Write to http error and print to stderr.
		http.Error(w, err.Error(), 400)
		log.Printf("Error: Failed to load environment variables %v", err)
		return
	}
	// write the jobs
	w.Write([]byte(jobsToString))
}
package getJobsForCandidatesHandler

import (
	"github.com/jasonjchu/bread/app/models/candidate"
	"github.com/jasonjchu/bread/app/models/job"
	"github.com/jasonjchu/bread/app/utils"
	"net/http"
	"strconv"
)

const RouteURL string = "/candidate"

func Handler(w http.ResponseWriter, r *http.Request) {
	cid := r.Header.Get("candidate_id")
	candidateId, err := strconv.Atoi(cid)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	r.ParseForm()
	tags := r.Form["tag_ids"]

	// retrieve the limit if present ow default to 200
	limit := r.URL.Query().Get("limit")
	jobLimit, err := strconv.Atoi(limit)
	if err != nil {
		jobLimit = 200
	}

	jobs, err := job.GetJobsByCidNotSeen(candidate.Id(candidateId), jobLimit, tags)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write([]byte(utils.ToJson(jobs)))
}

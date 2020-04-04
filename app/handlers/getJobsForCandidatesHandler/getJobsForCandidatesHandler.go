package getJobsForCandidatesHandler

import (
	"github.com/jasonjchu/bread/app/models/candidate"
	"github.com/jasonjchu/bread/app/models/job"
	"github.com/jasonjchu/bread/app/utils"
	"net/http"
	"strconv"
	"strings"
)

const RouteURL string = "/candidate"

func Handler(w http.ResponseWriter, r *http.Request) {
	cid := r.Header.Get("user_id")
	candidateId, err := strconv.Atoi(cid)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	tid := r.Header.Get("tag_ids")
	var tags []string
	// need to remove the brackets given an array
	if len(tid) > 2 {
		tags = strings.Split(tid[1:len(tid)-1], ",")
	}

	// retrieve the limit if present ow default to 200
	limit := r.Header.Get("job_limit")
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

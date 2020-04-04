package getCandidatesForJobHandler

import (
	"github.com/jasonjchu/bread/app/models/candidate"
	"github.com/jasonjchu/bread/app/utils"
	"net/http"
	"strconv"
)

const RouteURL string = "/candidates-for-job"

func Handler(w http.ResponseWriter, r *http.Request) {
	jobId := r.Header.Get("job_id")

	limit := r.Header.Get("candidate_limit")
	candidateLimit, err := strconv.Atoi(limit)
	if err != nil {
		candidateLimit = 200
	}

	candidates, err := candidate.GetCandidatesByJidLiked(jobId, candidateLimit)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write([]byte(utils.ToJson(candidates)))
}


package getCandidatesForJobHandler

import (
	"github.com/go-chi/chi"
	"github.com/jasonjchu/bread/app/models/candidate"
	"github.com/jasonjchu/bread/app/utils"
	"net/http"
	"strconv"
)

const RouteURL string = "/candidates"

func Handler(w http.ResponseWriter, r *http.Request) {
	jobId := chi.URLParam(r, "job_id")

	limit := r.URL.Query().Get("limit")
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

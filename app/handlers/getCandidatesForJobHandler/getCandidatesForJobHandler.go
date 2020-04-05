package getCandidatesForJobHandler

import (
	"github.com/go-chi/chi"
	"github.com/jasonjchu/bread/app/models/candidate"
	"github.com/jasonjchu/bread/app/utils"
	"net/http"
)

const RouteURL string = "/candidates"
const candidateLimit int = 200

func Handler(w http.ResponseWriter, r *http.Request) {
	jobId := chi.URLParam(r, "job_id")
	candidates, err := candidate.GetCandidatesByJidLiked(jobId, candidateLimit)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write([]byte(utils.ToJson(candidates)))
}

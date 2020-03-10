package getCandidatePoolHandler

import (
	"github.com/go-chi/chi"
	"github.com/jasonjchu/bread/app/models/candidate"
	"github.com/jasonjchu/bread/app/utils"
	"net/http"
	"strconv"
)

const RouteURL string = "/{id}"
const candidateLimit int = 200

func Handler(w http.ResponseWriter, r *http.Request) {
	jobId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	candidates, err := candidate.GetCandidatesByJid(candidate.Id(jobId), candidateLimit)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write([]byte(utils.ToJson(candidates)))
}


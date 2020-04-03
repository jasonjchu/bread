package candidateDislikesJobHandler

import (
	"github.com/go-chi/chi"
	"github.com/jasonjchu/bread/app/models/candidate"
	"github.com/jasonjchu/bread/app/models/job"
	"github.com/jasonjchu/bread/app/services/candidateDislikesJobService"
	"net/http"
	"strconv"
)

const RouteURL string = "/dislike"

func Handler(w http.ResponseWriter, r *http.Request) {
	cid := r.Header.Get("user_id")
	candidateId, err := strconv.Atoi(cid)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	jid := chi.URLParam(r, "job_id")
	req := candidateDislikesJobService.Request{
		CandidateId: candidate.Id(candidateId),
		JobId:       job.Id(jid),
	}
	err = candidateDislikesJobService.Exec(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(""))
}

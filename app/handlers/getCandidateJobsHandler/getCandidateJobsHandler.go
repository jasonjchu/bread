package getCandidateJobsHandler

import (
	"github.com/go-chi/chi"
	"github.com/jasonjchu/bread/app/models/job"
	"github.com/jasonjchu/bread/app/utils"
	"net/http"
)

const RouteURL string = "/{id}"
const jobLimit int = 200

func Handler(w http.ResponseWriter, r *http.Request) {
	candidateId := chi.URLParam(r, "id")
	jobs, err := job.GetJobsByCid(job.Id(candidateId), jobLimit)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write([]byte(utils.ToJson(jobs)))
}

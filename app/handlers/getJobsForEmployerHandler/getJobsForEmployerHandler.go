package getJobsForEmployerHandler

import (
	"github.com/jasonjchu/bread/app/models/employer"
	"github.com/jasonjchu/bread/app/services/employerJobsService"
	"github.com/jasonjchu/bread/app/utils"
	"net/http"
	"strconv"
)

const RouteURL string = "/"

func Handler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.Header.Get("user_id"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	req := employerJobsService.Request{Id: employer.Id(id)}
	jobs, err := employerJobsService.Exec(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(utils.ToJson(jobs)))
}

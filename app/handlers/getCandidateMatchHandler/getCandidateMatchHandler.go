package getCandidateMatchHandler

import (
	"github.com/jasonjchu/bread/app/models/candidate"
	"github.com/jasonjchu/bread/app/services/candidateMatchesService"
	"github.com/jasonjchu/bread/app/utils"
	"net/http"
	"strconv"
)

const RouteURL string = "/matches"

func Handler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.Header.Get("user_id"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	req := candidateMatchesService.Request{Id: candidate.Id(id)}
	matches, err := candidateMatchesService.Exec(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(utils.ToJson(matches)))
}

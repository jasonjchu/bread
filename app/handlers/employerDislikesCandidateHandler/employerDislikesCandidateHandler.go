package employerDislikesCandidateHandler

import (
	"encoding/json"
	"github.com/jasonjchu/bread/app/services/employerDislikesCandidateService"
	"net/http"
)

const RouteURL string = "/dislike"

func Handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req employerDislikesCandidateService.Request
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = employerDislikesCandidateService.Exec(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(""))
}

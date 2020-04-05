package employerLikesCandidateHandler

import (
	"encoding/json"
	"github.com/jasonjchu/bread/app/services/employerLikesCandidateService"
	"net/http"
)

const RouteURL string = "/like"

func Handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req employerLikesCandidateService.Request
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = employerLikesCandidateService.Exec(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(""))
}

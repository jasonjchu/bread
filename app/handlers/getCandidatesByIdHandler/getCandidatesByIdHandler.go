package getCandidatesByIdHandler

import (
	"github.com/go-chi/chi"
	"github.com/jasonjchu/bread/app/models/candidate"
	"github.com/jasonjchu/bread/app/utils"
	"net/http"
	"strconv"
)

const RouteURL string = "/{id}"

func Handler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	canId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	theCandidate, err := candidate.GetCandidateById(candidate.Id(canId))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write([]byte(utils.ToJson(theCandidate)))
}

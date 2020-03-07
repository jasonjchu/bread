package getEmployerHandler

import (
	"github.com/go-chi/chi"
	"github.com/jasonjchu/bread/app/models/employer"
	"github.com/jasonjchu/bread/app/utils"
	"net/http"
	"strconv"
)

const RouteURL string = "/{employer_id}"

func Handler(w http.ResponseWriter, r *http.Request) {
	employerIdStr := chi.URLParam(r, "employer_id")
	employerId, err := strconv.Atoi(employerIdStr)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	res, err := employer.GetEmployerById(employer.Id(employerId))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write([]byte(utils.ToJson(res)))
}

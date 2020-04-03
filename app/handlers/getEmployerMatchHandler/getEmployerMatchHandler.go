package getEmployerMatchHandler

import (
	"github.com/jasonjchu/bread/app/models/employer"
	"github.com/jasonjchu/bread/app/services/employerMatchesService"
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
	req := employerMatchesService.Request{Id: employer.Id(id)}
	matches, err := employerMatchesService.Exec(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(utils.ToJson(matches)))
}

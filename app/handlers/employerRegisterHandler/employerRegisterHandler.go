package employerRegisterHandler

import (
	"encoding/json"
	"github.com/jasonjchu/bread/app/services/employerRegisterService"
	"github.com/jasonjchu/bread/app/utils"
	"net/http"
)

const RouteURL string = "/register"

func Handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req employerRegisterService.Request
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	res, err := employerRegisterService.Exec(req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write([]byte(utils.ToJson(res)))
}

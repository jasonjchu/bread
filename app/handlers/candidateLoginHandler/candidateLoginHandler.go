package candidateLoginHandler

import (
	"encoding/json"
	"github.com/jasonjchu/bread/app/models/account"
	"github.com/jasonjchu/bread/app/services/candidateLoginService"
	"github.com/jasonjchu/bread/app/utils"
	"net/http"
)

const RouteURL string = "/login"

func Handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req candidateLoginService.Request
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	res, err := candidateLoginService.Exec(req)
	if err != nil {
		switch err.(type) {
		case account.InvalidPasswordError:
			http.Error(w, err.Error(), 401)
		default:
			http.Error(w, err.Error(), 400)
		}
		return
	}
	w.Write([]byte(utils.ToJson(res)))
}

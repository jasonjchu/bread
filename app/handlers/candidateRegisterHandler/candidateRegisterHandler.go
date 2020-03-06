package candidateRegisterHandler

import (
	"encoding/json"
	"github.com/jasonjchu/bread/app/models/account"
	"github.com/jasonjchu/bread/app/models/candidate"
	"github.com/jasonjchu/bread/app/utils"
	"net/http"
	"time"
)

const RouteURL string = "/register"

func Handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Name string `json:"name"`
		Program string `json:"program"`
		GradDate time.Time `json:"grad_date"`
		Description string `json:"description"`
	}
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	accountId, err := account.CreateAccount(req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = candidate.CreateCandidate(req.Name, req.Program, req.GradDate, req.Description, accountId)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	res := struct {
		AccountId account.Id `json:"account_id"`
	} {
		AccountId: accountId,
	}
	w.Write([]byte(utils.ToJson(res)))
}
package candidateRegisterService

import (
	"github.com/jasonjchu/bread/app/models/account"
	"github.com/jasonjchu/bread/app/models/candidate"
	"time"
)

type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name string `json:"name"`
	Program string `json:"program"`
	GradDate time.Time `json:"grad_date"`
	Description string `json:"description"`
}

type Response struct {
	AccountId account.Id `json:"account_id"`
}

func Exec(req Request) (*Response, error) {
	accountId, err := account.CreateAccount(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	err = candidate.CreateCandidate(req.Name, req.Program, req.GradDate, req.Description, accountId)
	if err != nil {
		return nil, err
	}
	res := Response{AccountId: accountId}
	return &res, nil
}

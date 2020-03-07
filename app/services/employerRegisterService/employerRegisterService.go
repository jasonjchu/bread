package employerRegisterService

import (
	"github.com/jasonjchu/bread/app/models/account"
	"github.com/jasonjchu/bread/app/models/employer"
)

type Request struct {
	Username account.Username `json:"username"`
	Password account.Password `json:"password"`
	Name     string           `json:"name"`
	WorksAt  int              `json:"worksAt"`
}

type Response struct {
	AccountId account.Id `json:"account_id"`
}

func Exec(req Request) (*Response, error) {
	accountId, err := account.CreateAccount(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	name := employer.Name(req.Name)
	worksAt := employer.WorksAt(req.WorksAt)
	err = employer.CreateEmployer(name, worksAt, accountId)
	if err != nil {
		return nil, err
	}
	res := Response{AccountId: accountId}
	return &res, nil
}

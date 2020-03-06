package employerLoginService

import (
	"github.com/jasonjchu/bread/app/models/account"
	"github.com/jasonjchu/bread/app/models/employer"
)

type Request struct {
	Username account.Username `json:"username"`
	Password account.Password `json:"password"`
}

type Response struct {
	Id int `json:"id"`
}

func Exec(req Request) (*Response, error) {
	acc, err := account.GetAccountByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	err = account.VerifyPassword(acc.Password, req.Password)
	if err != nil {
		return nil, err
	}
	// Ensure that we have an employer account
	_, err = employer.GetEmployerById(employer.Id(int(acc.Id)))
	if err != nil {
		return nil, err
	}
	res := Response{
		Id: int(acc.Id),
	}
	return &res, nil
}

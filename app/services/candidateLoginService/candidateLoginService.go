package candidateLoginService

import (
	"github.com/jasonjchu/bread/app/models/account"
	"github.com/jasonjchu/bread/app/models/candidate"
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
	// Ensure that we have a candidate account
	_, err = candidate.GetCandidateById(candidate.Id(int(acc.Id)))
	if err != nil {
		return nil, err
	}
	res := Response{
		Id: int(acc.Id),
	}
	return &res, nil
}

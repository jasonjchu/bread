package getCompaniesHandler

import (
	"github.com/jasonjchu/bread/app/models/company"
	"github.com/jasonjchu/bread/app/utils"
	"net/http"
)

const RouteURL string = "/"

type ResponseWrapper struct {
	Companies company.Companies `json:"companies"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	companies, err := company.GetCompaniesWithName(name)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	response := ResponseWrapper{Companies: companies}
	w.Write([]byte(utils.ToJson(response)))
}

package getEmployersHandler

import (
	"github.com/jasonjchu/bread/app/utils"
	"net/http"
)
const RouteURL string = "/"

func Handler(w http.ResponseWriter, r *http.Request) {
	// TODO: Get the context from r to specify the number of jobs / country code / has expired / etc

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write([]byte(utils.ToJson(employers)))
}

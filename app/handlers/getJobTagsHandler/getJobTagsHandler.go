package getJobTagsHandler

import (
	"github.com/jasonjchu/bread/app/models/tag"
	"github.com/jasonjchu/bread/app/utils"
	"net/http"
)

const RouteURL string = "/tags"

func Handler(w http.ResponseWriter, r *http.Request) {
	// call function to get all tags
	tags, err := tag.GetAllJobTags()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(utils.ToJson(tags)))
}
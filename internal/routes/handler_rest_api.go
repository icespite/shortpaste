package routes

import (
	"github.com/timeforaninja/shortpaste/internal/routes/generic_rest"
	"github.com/timeforaninja/shortpaste/internal/types"
	"github.com/timeforaninja/shortpaste/internal/utils"
	"net/http"
	"strings"
)

func HandleRestAPI(app types.AppInf) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/v1")
		switch r.Method {
		case "GET":
			switch {
			case strings.HasPrefix(r.URL.Path, "/f/"):
				generic_rest.HandleGetFiles(w, r, app)
			case strings.HasPrefix(r.URL.Path, "/l/"):
				generic_rest.HandleGetLinks(w, r, app)
			case strings.HasPrefix(r.URL.Path, "/t/"):
				generic_rest.HandleGetTexts(w, r, app)
			default:
				utils.OnNotFound(w, "No such endpoint")
			}

		case "POST":
			switch {
			case strings.HasPrefix(r.URL.Path, "/f/"):
				generic_rest.HandleCreateFile(w, r, app)
			case strings.HasPrefix(r.URL.Path, "/l/"):
				generic_rest.HandleCreateLink(w, r, app)
			case strings.HasPrefix(r.URL.Path, "/t/"):
				generic_rest.HandleCreateText(w, r, app)
			default:
				utils.OnNotFound(w, "No such endpoint")
			}

		default:
			utils.OnMethodNotAllowed(w, "Only GET and POST Supported")
		}
	}
}

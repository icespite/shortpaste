package handler_templates

import (
	"fmt"
	"github.com/timeforaninja/shortpaste/internal/routes/handler_templates/file_templates"
	"github.com/timeforaninja/shortpaste/internal/types"
	"github.com/timeforaninja/shortpaste/internal/utils"
	"net/http"
	"strings"
)

func ResolveShortLink(app types.AppInf) http.HandlerFunc {
	ddb := app.GetDataDB()
	link307Redirect := app.ShouldLink307Redirect()
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/l/"), "/")
		if id == "" {
			utils.OnNotFound(w, "No ID found in request")
			return
		}
		var link types.Link
		err := ddb.First(&link, "id = ?", id).Error
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Link for `%s` not found!\n", id)
			return
		}

		if link307Redirect {
			http.Redirect(w, r, link.Link, http.StatusTemporaryRedirect)
		} else {
			t, err := file_templates.LoadTemplate("file_templates/link.html")
			if err != nil {
				utils.OnServerError(w, err, "failed to parse template")
				return
			}
			t.Execute(w, link)
		}

		link.HitCount += 1
		ddb.Save(&link)
	}
}

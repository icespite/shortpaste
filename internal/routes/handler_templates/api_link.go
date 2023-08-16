package handler_templates

import (
	"fmt"
	"github.com/timeforaninja/shortpaste/internal/routes/handler_templates/file_templates"
	"github.com/timeforaninja/shortpaste/internal/types"
	"github.com/timeforaninja/shortpaste/internal/utils"
	"net/http"
	"strings"
	"time"
)

func ResolveShortLink(app types.AppInf) http.HandlerFunc {
	ddb := app.GetDataDB()
	link307Redirect := app.ShouldLink307Redirect()
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/l/"), "/")
		if id == "" {
			utils.LogIfDebug("No ID found")
			utils.OnNotFound(w, "No ID found in request")
			return
		}
		var link types.Link
		err := ddb.First(&link, "id = ?", id).Error
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			utils.LogIfDebug("Failed to find Link object: %s", err)
			fmt.Fprintf(w, "Link for `%s` not found!\n", id)
			return
		}

		if link307Redirect {
			http.Redirect(w, r, link.Link, http.StatusTemporaryRedirect)
		} else {
			t, err := file_templates.LoadTemplate("link.html")
			if err != nil {
				utils.LogIfDebug("Failed to insert into template: %s\nobject: %s", err, link.Link)
				utils.OnServerError(w, err, "failed to parse template")
				return
			}
			t.Execute(w, link)
		}

		link.HitCount += 1
		link.LastHit = time.Now().Unix()
		ddb.Save(&link)
	}
}

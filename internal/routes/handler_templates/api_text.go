package handler_templates

import (
	"fmt"
	"github.com/timeforaninja/shortpaste/internal/routes/handler_templates/file_templates"
	"github.com/timeforaninja/shortpaste/internal/types"
	"github.com/timeforaninja/shortpaste/internal/utils"
	"html"
	"net/http"
	"path"
	"strings"
	"time"
)

func ResolveShortText(app types.AppInf) http.HandlerFunc {
	ddb := app.GetDataDB()
	fdb := app.GetFileDB()
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/t/"), "/")
		if id == "" {
			utils.OnNotFound(w, "No ID found in request")
			return
		}

		// read Text from db
		var text types.Text
		err := ddb.First(&text, "id = ?", id).Error
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Text for `%s` not found!\n", id)
			return
		}

		subPath := path.Join("texts", text.ID+"."+text.Type)
		// check if download direct query is set
		if _, ok := r.URL.Query()["download"]; ok {
			w.Header().Set("Content-Disposition", "attachment; filename="+text.ID+"."+text.Type)
			fdb.ServeFile(w, r, subPath)

			// update the access counter
			text.HitCount += 1
			text.LastHit = time.Now().Unix()
			ddb.Save(&text)
			return
		}

		// load template and embed data
		t, err := file_templates.LoadTemplate("text.html")
		if err != nil {
			utils.OnServerError(w, err, "failed to parse template")
			return
		}
		textContent, err := fdb.Read(subPath)
		if err != nil {
			utils.OnServerError(w, err, "failed to read text")
			return
		}

		// should highlight?
		var highlight string
		if text.NoHighlight {
			highlight = "language-plaintext"
		}

		// pack data into struct, embed into template and send
		data := struct {
			ID    string
			Class string
			Text  string
		}{
			ID:    text.ID,
			Class: highlight,
			Text:  html.EscapeString(string(textContent)),
		}
		t.Execute(w, data)

		// update the access counter
		text.HitCount += 1
		text.LastHit = time.Now().Unix()
		ddb.Save(&text)
		return
	}
}

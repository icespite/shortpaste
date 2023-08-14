package handler_templates

import (
	"fmt"
	"github.com/timeforaninja/shortpaste/internal/routes/handler_templates/file_templates"
	"github.com/timeforaninja/shortpaste/internal/types"
	"github.com/timeforaninja/shortpaste/internal/utils"
	"net/http"
	"path"
	"strings"
)

func ResolveShortFile(app types.AppInf) http.HandlerFunc {
	ddb := app.GetDataDB()
	fdb := app.GetFileDB()
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/f/"), "/")
		if id == "" {
			utils.OnNotFound(w, "No ID found in request")
			return
		}

		// read File from db
		var file types.File
		err := ddb.First(&file, "id = ?", id).Error
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Link for `%s` not found!\n", id)
			return
		}

		subPath := path.Join("files", file.ID, file.Name)

		// check if download direct query is set
		if _, ok := r.URL.Query()["download"]; ok {
			w.Header().Set("Content-Disposition", "attachment; filename="+file.Name)
			fdb.ServeFile(w, r, subPath)

			// update the access counter
			file.DownloadCount += 1
			ddb.Save(&file)
			return
		}

		// load template and embed data
		t, err := file_templates.LoadTemplate("file_templates/files.html")
		if err != nil {
			utils.OnServerError(w, err, "failed to parse template")
			return
		}
		fi, err := fdb.Stat(subPath)
		if err != nil {
			utils.OnServerError(w, err, "failed to get file size")
			return
		}

		// pack data into struct, embed into template and send
		data := struct {
			Name  string
			Src   string
			Image bool
			Size  string
		}{
			Name:  file.Name,
			Src:   "/f/" + id + "?download",
			Image: strings.HasPrefix(file.MIME, "image/"),
			Size:  utils.IECFormat(fi.Size()),
		}
		t.Execute(w, data)

		// update the access counter
		file.HitCount += 1
		ddb.Save(&file)
	}
}

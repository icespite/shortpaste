package generic_rest

import (
	"encoding/json"
	"github.com/timeforaninja/shortpaste/internal/types"
	"github.com/timeforaninja/shortpaste/internal/utils"
	"net/http"
	"path"
	"strings"
)

func HandleCreateText(w http.ResponseWriter, r *http.Request, app types.AppInf) {
	text := types.Text{}
	if err := json.NewDecoder(r.Body).Decode(&text); err != nil {
		utils.OnClientError(w, err, "check the input and try again")
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/l/")
	if id != "" && text.ID == "" {
		text.ID = id
	}
	if err := text.Validate(); err != nil {
		utils.OnClientError(w, err, "check the input and try again")
		return
	}

	if text.Type == "" {
		text.Type = "txt"
	}

	// save text on disk
	subPath := path.Join("texts", text.ID+"."+text.Type)
	err := app.GetFileDB().Write(subPath, text.Text)
	if err != nil {
		utils.OnServerError(w, err, "failed to write contents to disk")
		return
	}

	// save metadata in db
	err = app.GetDataDB().Create(&text).Error
	if err != nil {
		utils.OnServerError(w, err, "failed to create DB entry")
		return
	}

	// respond with success status
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "created"})
}

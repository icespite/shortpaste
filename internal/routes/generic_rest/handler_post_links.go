package generic_rest

import (
	"encoding/json"
	"github.com/timeforaninja/shortpaste/internal/types"
	"github.com/timeforaninja/shortpaste/internal/utils"
	"net/http"
	"strings"
)

func HandleCreateLink(w http.ResponseWriter, r *http.Request, app types.AppInf) {
	link := types.Link{}
	if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
		utils.OnClientError(w, err, "check the input and try again")
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/l/")
	if id != "" && link.ID == "" {
		link.ID = id
	}
	if err := link.Validate(); err != nil {
		utils.OnClientError(w, err, "check the input and try again")
		return
	}
	err := app.GetDataDB().Create(&link).Error
	if err != nil {
		utils.OnServerError(w, err, "failed to create DB entry")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "created"})
}

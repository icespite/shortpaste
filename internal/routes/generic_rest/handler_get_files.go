package generic_rest

import (
	"encoding/json"
	"github.com/timeforaninja/shortpaste/internal/types"
	"net/http"
)

func HandleGetFiles(w http.ResponseWriter, r *http.Request, app types.AppInf) {
	var files []types.File
	app.GetDataDB().Find(&files)
	json.NewEncoder(w).Encode(map[string][]types.File{"files": files})
}

package generic_rest

import (
	"encoding/json"
	"github.com/timeforaninja/shortpaste/internal/types"
	"net/http"
)

func HandleGetLinks(w http.ResponseWriter, r *http.Request, app types.AppInf) {
	var links []types.Link
	app.GetDataDB().Find(&links)
	json.NewEncoder(w).Encode(map[string][]types.Link{"links": links})
}

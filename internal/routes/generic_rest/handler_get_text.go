package generic_rest

import (
	"encoding/json"
	"github.com/timeforaninja/shortpaste/internal/types"
	"net/http"
)

func HandleGetTexts(w http.ResponseWriter, r *http.Request, app types.AppInf) {
	var texts []types.Text
	app.GetDataDB().Find(&texts)
	json.NewEncoder(w).Encode(map[string][]types.Text{"texts": texts})
}

package middleware

import (
	"github.com/timeforaninja/shortpaste/internal/types"
	"github.com/timeforaninja/shortpaste/internal/utils"
	"net/http"
)

func EnforceAuth(app types.AppInf, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Fetch Authorization
		username, password, ok := r.BasicAuth()
		if !ok {
			// no Auth Present -> request it
			utils.OnNotAuthorizedError(w)
			return
		}

		shouldU, shouldPW := app.GetDataDB().GetUserPass()
		validLogin := utils.CompareAuth(username, password, shouldU, shouldPW)

		if !validLogin {
			// invalid login -> re-request auth
			utils.OnNotAuthorizedError(w)
			return
		}

		next.ServeHTTP(w, r)
	}
}

package middleware

import (
	"github.com/timeforaninja/shortpaste/internal/utils"
	"net/http"
)

func EnforceGet(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
			return
		}

		utils.OnMethodNotAllowed(w, "Only GET allowed")
	}
}

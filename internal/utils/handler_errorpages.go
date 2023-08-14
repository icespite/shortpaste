package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func OnClientError(w http.ResponseWriter, err error, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("%v", err), "message": msg})
}

func OnNotFound(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": msg})
}

func OnServerError(w http.ResponseWriter, err error, msg string) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("%v", err), "message": msg})
}

func OnNotAuthorizedError(w http.ResponseWriter) {
	// If the Authentication header is not present, is invalid, or the
	// username or password is wrong, then set a WWW-Authenticate
	// header to inform the client that we expect them to use basic
	// authentication and send a 401 Unauthorized response.
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func OnMethodNotAllowed(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(map[string]string{"message": msg})
}

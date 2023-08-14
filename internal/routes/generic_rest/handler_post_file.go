package generic_rest

import (
	"encoding/json"
	"fmt"
	"github.com/timeforaninja/shortpaste/internal/types"
	"github.com/timeforaninja/shortpaste/internal/utils"
	"net/http"
	"path"
	"strings"
)

func HandleCreateFile(w http.ResponseWriter, r *http.Request, app types.AppInf) {
	file := types.File{}
	file.ID = strings.TrimPrefix(r.URL.Path, "/f/")
	if err := file.Validate(); err != nil {
		utils.OnClientError(w, err, "check the input and try again")
		return
	}

	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and headers
	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		utils.OnClientError(w, err, "failed to retrieve file, check if the upload completed")
		return
	}
	defer uploadedFile.Close()

	file.Name = handler.Filename
	file.MIME = handler.Header["Content-Type"][0]

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create folder & file + copy content from request body
	subPath := path.Join("files", file.ID, file.Name)
	err = app.GetFileDB().WriteStream(subPath, uploadedFile)
	if err != nil {
		utils.OnServerError(w, err, "failed writing file to disk")
		return
	}

	// save metadata in db
	err = app.GetDataDB().Create(&file).Error
	if err != nil {
		utils.OnServerError(w, err, "failed to create DB entry")
		return
	}

	// respond with success status
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "created"})
}

package types

// File struct for saving the file uploads /f/.
type File struct {
	BaseItem
	Name          string `json:"name"`
	MIME          string `validate:"isdefault"`
	DownloadCount int64  `json:"downloadcount" validate:"isdefault"`
}

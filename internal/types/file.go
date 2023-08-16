package types

// File struct for saving the file uploads /f/.
type File struct {
	BaseItem
	ID            string `gorm:"primaryKey" json:"id" validate:"required,min=3,max=32,alphanumunicode"`
	Name          string `json:"name"`
	MIME          string `validate:"isdefault"`
	DownloadCount int64  `json:"downloadcount" validate:"isdefault"`
}

package types

import (
	"gorm.io/gorm"
)

// File struct for saving the file uploads /f/.
type File struct {
	ID            string `gorm:"primaryKey" json:"id" validate:"required,min=3,max=32,alphanumunicode"`
	Name          string `json:"name"`
	MIME          string `validate:"isdefault"`
	HitCount      int64  `json:"hitcount" validate:"isdefault"`
	DownloadCount int64  `json:"downloadcount" validate:"isdefault"`
	gorm.Model
}

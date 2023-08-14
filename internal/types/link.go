package types

import (
	"gorm.io/gorm"
)

// Link struct for saving the Redirect Links /l/.
type Link struct {
	ID       string `gorm:"primaryKey" json:"id" validate:"required,min=3,max=32,alphanumunicode"`
	Link     string `json:"link" validate:"required,url"`
	HitCount int64  `json:"hitcount" validate:"isdefault"`
	gorm.Model
}

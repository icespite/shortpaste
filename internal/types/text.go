package types

import (
	"gorm.io/gorm"
)

// Text struct for saving the text pastes /t/.
type Text struct {
	ID          string `gorm:"primaryKey" json:"id" validate:"required,min=3,max=32,alphanumunicode"`
	Type        string `validate:"omitempty,oneof=txt md" json:"type"`
	Text        string `gorm:"-" json:"text,omitempty"`
	NoHighlight bool   `json:"nohighlight"`
	HitCount    int64  `json:"hitcount" validate:"isdefault"`
	gorm.Model
}

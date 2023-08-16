package types

import "gorm.io/gorm"

type BaseItem struct {
	ID       string `gorm:"primaryKey" json:"id" validate:"required,min=3,max=32,alphanumunicode"`
	HitCount int64  `json:"hitcount" validate:"isdefault"`
	LastHit  int64  `json:"lasthit" validate:"min=0"`
	gorm.Model
}

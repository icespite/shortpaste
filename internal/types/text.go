package types

// Text struct for saving the text pastes /t/.
type Text struct {
	BaseItem
	Type        string `validate:"omitempty,oneof=txt md" json:"type"`
	Text        string `gorm:"-" json:"text,omitempty"`
	NoHighlight bool   `json:"nohighlight"`
}

package types

// Link struct for saving the Redirect Links /l/.
type Link struct {
	BaseItem
	Link string `json:"link" validate:"required,url"`
}

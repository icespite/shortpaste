package types

import (
	"gopkg.in/go-playground/validator.v9"
)

func (link *Link) Validate() error {
	return validator.New().Struct(link)
}

func (file *File) Validate() error {
	return validator.New().Struct(file)
}

func (text *Text) Validate() error {
	return validator.New().Struct(text)
}

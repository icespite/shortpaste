package data_db

import (
	"gopkg.in/go-playground/validator.v9"
)

func (link *types.Link) Validate() error {
	return validator.New().Struct(link)
}

func (file *types.File) Validate() error {
	return validator.New().Struct(file)
}

func (text *types.Text) Validate() error {
	return validator.New().Struct(text)
}

package file_templates

import (
	"embed"
	"text/template"
)

// Embed html file_templates with code
//
//go:embed *.html
var templateFS embed.FS

func LoadTemplate(path string) (*template.Template, error) {
	t, err := template.ParseFS(templateFS, path)
	return t, err
}

package file_templates

import (
	"embed"
	"github.com/timeforaninja/shortpaste/internal/utils"
	"text/template"
)

// Embed html file_templates with code
//
//go:embed *.html
var templateFS embed.FS

func LoadTemplate(path string) (*template.Template, error) {
	t, err := template.ParseFS(templateFS, path)
	utils.LogIfDebug("Error loading Template %s", err)
	return t, err
}

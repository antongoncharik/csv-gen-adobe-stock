package handlers

import (
	"html/template"
	"log"
	"path/filepath"
)

var tmpl *template.Template

func LoadTemplates(templatesDir string) {
	tmpl = template.Must(template.ParseGlob(filepath.Join(templatesDir, "*.html")))
	log.Println("Templates loaded")
}

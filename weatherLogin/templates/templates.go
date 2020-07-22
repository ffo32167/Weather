package templates

import (
	"html/template"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// Templates шаблоны страниц
var Templates *template.Template

// Initialize инициализирует шаблоны
func Initialize(path string) {
	templatesPath := filepath.Join(path, `static`, `*.html`)
	t, err := template.ParseGlob(templatesPath)
	if err != nil {
		logrus.Fatal("template.ParseGlob err:", err)
	}
	Templates = template.Must(t, err)
}

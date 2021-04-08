// @path     template.go
// @file     template.go
// @author   Marc Proux <marc.proux@outlook.fr>
// @date     Thu, 08 Apr 2021 14:40:23 GMT
// @modified Thu, 08 Apr 2021 16:21:28 GMT

package config

import (
	"os"
	"path/filepath"
	"text/template"
)

// ===== STRUCTURE ================================================================================

type Template struct {
	Func template.FuncMap
}

// ===== EXTERNAL =================================================================================

// ProcessFile va lancer la procéder à l'élaboration d'un fichier template.
func (t *Template) ProcessFile(path string) (outputPath string, err error) {
	tmpl, err := template.New(filepath.Base(path)).Funcs(t.Func).ParseFiles(path)
	if err != nil {
		return "", err
	}

	outputPath = filepath.Join("/tmp", filepath.Base(path))
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer outputFile.Close()

	err = tmpl.Execute(outputFile, nil)
	if err != nil {
		return "", err
	}

	return outputPath, nil
}

// ===== INTERNAL =================================================================================

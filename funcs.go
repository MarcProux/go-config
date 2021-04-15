// @path     funcs.go
// @file     funcs.go
// @author   Marc Proux <marc.proux@outlook.fr>
// @date     Thu, 15 Apr 2021 08:09:42 GMT
// @modified Thu, 15 Apr 2021 08:14:48 GMT

package config

import (
	"os"
	"text/template"

	"github.com/Masterminds/sprig"
)

// ===== STRUCTURE ================================================================================

// ===== EXTERNAL =================================================================================

// ===== INTERNAL =================================================================================

func funcMap() template.FuncMap {
	r := template.FuncMap{
		"cFileContents": fileContents,
	}

	// Add Sprig functions
	for k, v := range sprig.FuncMap() {
		r[k] = v
	}

	return r
}

func fileContents() func(string) (string, error) {
	return func(s string) (string, error) {
		if s == "" {
			return "", nil
		}
		contents, err := os.ReadFile(s)
		if err != nil {
			return "", err
		}
		return string(contents), nil
	}
}

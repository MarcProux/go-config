// @path     funcs.go
// @file     funcs.go
// @author   Marc Proux <marc.proux@outlook.fr>
// @date     Thu, 15 Apr 2021 08:09:42 GMT
// @modified Tue, 27 Jul 2021 14:36:36 GMT

package config

import (
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

// ===== STRUCTURE ================================================================================

// ===== EXTERNAL =================================================================================

// ===== INTERNAL =================================================================================

func funcMap() template.FuncMap {
	r := sprig.TxtFuncMap()

	l := template.FuncMap{
		"fileContents": fileContents(),
	}

	for k, v := range l {
		if _, ok := r[k]; ok {
			k = "c_" + k
		}
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

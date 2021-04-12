// @path     item.go
// @file     item.go
// @author   Marc Proux <marc.proux@outlook.fr>
// @date     Mon, 12 Apr 2021 14:11:36 GMT
// @modified Mon, 12 Apr 2021 22:07:54 GMT

package config

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/mcuadros/go-lookup"
	"gopkg.in/yaml.v2"
)

// ===== STRUCTURE ================================================================================

type Item struct {
	Object   interface{}
	Path     string
	Template Template
	Type     ConfigType
}

// ===== EXTERNAL =================================================================================

func (i *Item) Key(path string, delimiter ...string) (value interface{}, err error) {
	if len(delimiter) == 0 {
		delimiter = make([]string, 1)
		delimiter[0] = "::"
	}
	value, err = lookup.Lookup(i.Object, strings.Split(path, delimiter[0])...)
	return value, err
}

// ===== INTERNAL =================================================================================

func (i *Item) loadConfig() (err error) {
	if _, err = os.Stat(i.Path); os.IsNotExist(err) {
		return err
	}

	outputFile, err := i.Template.ProcessFile(i.Path)
	if err != nil {
		return err
	}

	err = nil
	switch i.Type {

	case TypeHCL:
		err = hclsimple.DecodeFile(outputFile, nil, i.Object)

	case TypeJSON:
		data, _ := os.ReadFile(outputFile)
		err = json.Unmarshal(data, i.Object)

	case TypeYAML:
		data, _ := os.ReadFile(outputFile)
		err = yaml.Unmarshal(data, i.Object)

	}

	return err
}

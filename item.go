// @path     item.go
// @file     item.go
// @author   Marc Proux <marc.proux@outlook.fr>
// @date     Mon, 12 Apr 2021 14:11:36 GMT
// @modified Tue, 13 Apr 2021 16:29:03 GMT

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

func (i *Item) Key(path string, delimiter ...byte) (value interface{}, err error) {
	if path == "" {
		return i.Object, nil
	}

	if len(delimiter) == 0 {
		delimiter = make([]byte, 1)
		delimiter[0] = ':'
	}
	lValue, err := lookup.Lookup(i.Object, strings.Split(path, string(delimiter[0]))...)
	return lValue.Interface(), err
}

func (i *Item) KeyI(path string, delimiter ...byte) (interface{}, error) {
	if path == "" {
		return i.Object, nil
	}

	if len(delimiter) == 0 {
		delimiter = make([]byte, 1)
		delimiter[0] = ':'
	}
	lValue, err := lookup.LookupI(i.Object, strings.Split(path, string(delimiter[0]))...)
	return lValue.Interface(), err
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

// @path     examples/json/main.go
// @file     main.go
// @author   Marc Proux <marc.proux@outlook.fr>
// @date     Mon, 12 Apr 2021 21:26:33 GMT
// @modified Mon, 12 Apr 2021 22:08:13 GMT

package main

import (
	"fmt"

	"github.com/MarcProux/go-config"
)

// ===== STRUCTURE ================================================================================

type Config struct {
	Default string      `json:"default"`
	Field   string      `json:"field"`
	Block   ConfigBlock `json:"block"`
}

type ConfigBlock struct {
	Default string `json:"default"`
	Field   string `json:"field"`
}

// ===== EXTERNAL =================================================================================

// ===== INTERNAL =================================================================================

var def = &Config{
	Default: "default_string",
	Block: ConfigBlock{
		Default: "default_string",
	},
}

func main() {
	m, err := config.New()
	if err != nil {
		panic(err)
	}

	err = m.AddConfig(config.Config{
		Object: &def,
		Name:   "config",
		Path:   "config.json",
		Type:   config.TypeJSON,
	}, false)
	if err != nil {
		panic(err)
	}

	i, err := m.Get("config")
	if err != nil {
		panic(err)
	}

	v, err := i.Key("Block")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", v)
}

// @path     examples/hcl/main.go
// @file     main.go
// @author   Marc Proux <marc.proux@outlook.fr>
// @date     Mon, 12 Apr 2021 16:38:56 GMT
// @modified Mon, 12 Apr 2021 22:08:43 GMT

package main

import (
	"fmt"

	"github.com/MarcProux/go-config"
)

// ===== STRUCTURE ================================================================================

type Config struct {
	Default string      `hcl:"default,optional"`
	Field   string      `hcl:"field"`
	Block   ConfigBlock `hcl:"block,block"`
}

type ConfigBlock struct {
	Default string `hcl:"default,optional"`
	Field   string `hcl:"field"`
}

// ===== EXTERNAL =================================================================================

// ===== INTERNAL =================================================================================

var def = Config{
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
		Path:   "config.hcl",
		Type:   config.TypeHCL,
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

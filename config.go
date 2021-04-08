// @path     config.go
// @file     config.go
// @author   Marc Proux <marc.proux@outlook.fr>
// @date     Thu, 08 Apr 2021 13:22:29 GMT
// @modified Thu, 08 Apr 2021 16:24:29 GMT

package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/mcuadros/go-lookup"
)

// ===== STRUCTURE ================================================================================

type Config struct {
	Object   interface{}
	Template Template
}

type CManager struct {
	obj      interface{} // Configuration type structure
	template Template
}

// ===== EXTERNAL =================================================================================

func New(config Config) (*CManager, error) {
	c := &CManager{}

	if config.Object == nil {
		return nil, errors.New("field Object couldn't be nil")
	}
	if reflect.ValueOf(config.Object).Kind() != reflect.Ptr {
		return nil, errors.New("field Object has to be a pointer")
	}

	c.obj = config.Object
	c.template = config.Template

	return c, nil
}

// FromFile va lire et instancier la configuration depuis le chemin de fichier passé en paramètre.
func (c *CManager) FromFile(path string) (err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		return err
	}

	outputFile, err := c.template.ProcessFile(path)
	if err != nil {
		return err
	}
	err = hclsimple.DecodeFile(outputFile, nil, c.obj)
	return err
}

// GetKey permet d'explorer la configuration afin de récupérer une donnée suivant un chemin de
// clé séparé par un délimiteur (par défaut "::").
func (c *CManager) GetKey(key string, delimiter ...string) (value interface{}, err error) {
	if len(delimiter) == 0 {
		delimiter = make([]string, 1)
		delimiter[0] = "::"
	}
	value, err = lookup.LookupI(c.obj, strings.Split(key, delimiter[0])...)
	return value, err
}

// GetKeys permet d'explorer la configuration afin de récupérer des données suivant des chemins de
// clés séparé par un délimiteur (par défaut "::").
func (c *CManager) GetKeys(keys []string, delimiter ...string) (values []interface{}, err error) {
	for _, key := range keys {
		value, err := c.GetKey(key, delimiter...)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}

// Validate permet de vérifier la syntaxe du fichier de configuration selon les principes HCLv2.
func (c *CManager) Validate(path string) (err error) {
	// Vérifie que le fichier existe
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	// Résoud le template du fichier
	file, err := c.template.ProcessFile(path)
	if err != nil {
		return err
	}

	// Lit le fichier
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	_, diags := hclsyntax.ParseConfig(
		data, filepath.Join("/tmp", filepath.Base(path)), hcl.Pos{Line: 1, Column: 1},
	)
	if diags.HasErrors() {
		errors := ""
		for _, e := range diags.Errs() {
			errors = fmt.Sprintf("%s%s\n", errors, e)
		}
		return fmt.Errorf(errors)
	}

	return nil
}

// ===== INTERNAL =================================================================================

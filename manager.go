// @path     manager.go
// @file     manager.go
// @author   Marc Proux <marc.proux@outlook.fr>
// @date     Mon, 12 Apr 2021 14:11:17 GMT
// @modified Thu, 15 Apr 2021 08:15:09 GMT

package config

import (
	"errors"
	"fmt"
	"reflect"
	"text/template"
)

// ===== STRUCTURE ================================================================================

type Config struct {
	Name         string
	Object       interface{}
	Path         string
	TemplateFunc template.FuncMap
	Type         ConfigType
}

type Manager struct {
	configs map[string]*Item
}

// ===== EXTERNAL =================================================================================

func (m *Manager) AddConfig(config Config, override bool) (err error) {

	if err = m.check(&config); err != nil {
		return err
	}

	if _, ok := m.configs[config.Name]; ok && !override {
		return fmt.Errorf("%s already exists", config.Name)
	}

	t := Template{}
	t.Func = make(template.FuncMap)
	for k, v := range funcMap() {
		t.Func[k] = v
	}
	for k, v := range config.TemplateFunc {
		t.Func[k] = v
	}

	i := &Item{
		Object:   config.Object,
		Path:     config.Path,
		Template: t,
		Type:     config.Type,
	}
	err = i.loadConfig()
	if err != nil {
		return err
	}

	m.configs[config.Name] = i

	return nil
}

func (m *Manager) Get(name string) (item *Item, err error) {
	if _, ok := m.configs[name]; !ok {
		return nil, fmt.Errorf("%s does not exist in config", name)
	}
	return m.configs[name], nil
}

func Instance() *Manager {
	return manager
}

// ===== INTERNAL =================================================================================

var manager *Manager

func (m *Manager) check(config *Config) (err error) {
	if config.Object == nil {
		return errors.New("field Object couldn't be nil")
	}
	if reflect.ValueOf(config.Object).Kind() != reflect.Ptr {
		return errors.New("field Object must be a pointer")
	}

	if config.Path == "" {
		return errors.New("field Path couldn't be empty")
	}

	if config.Name == "" {
		config.Name = "config"
	}

	if config.TemplateFunc == nil {
		config.TemplateFunc = make(map[string]interface{})
	}

	if config.Type == "" {
		config.Type = TypeHCL
	}

	return nil
}

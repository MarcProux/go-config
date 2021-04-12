// @path     config.go
// @file     config.go
// @author   Marc Proux <marc.proux@outlook.fr>
// @date     Mon, 12 Apr 2021 14:10:26 GMT
// @modified Mon, 12 Apr 2021 15:02:58 GMT

package config

// ===== STRUCTURE ================================================================================

type ConfigType string

const (
	TypeHCL  ConfigType = "hcl"
	TypeJSON ConfigType = "json"
	TypeYAML ConfigType = "yaml"
)

// ===== EXTERNAL =================================================================================

func New() (m *Manager, err error) {
	m = &Manager{}

	m.configs = make(map[string]*Item)

	manager = m
	return m, nil
}

// ===== INTERNAL =================================================================================

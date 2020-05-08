package options

import (
	"todoList/conf/parse_yaml/conf"
)

// Shared Options
var (
	Options *AppOptions
)

// AppOptions defines specs for config
type AppOptions struct {
	Version string `yaml:"version"`
	GinService Service `yaml:"service"`

}

func NewAppOptions()error {
	return conf.Unmarshal(&Options)
}
type Service struct {
	Listen string `yaml:"listen"`
}

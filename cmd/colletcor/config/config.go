package config

import (
	"github.com/Zelayan/dts/pkg/store"
)

type Config struct {
	Default   DefaultOptions
	StoreType store.StoreType
}

type DefaultOptions struct {
	Listen string `yaml:"listen"`
}

// Valid 校验配置
func (c *Config) Valid() error {
	return nil
}

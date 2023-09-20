package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	yamlConfig = "yaml"
)

type StoreType = string

type Config struct {
	configFile string
	configType string
	data       []byte

	Default   DefaultOptions `yaml:"default"`
	StoreType StoreType      `yaml:"storeType"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) SetConfigType(in string) {
	c.configType = in
}
func (c *Config) SetConfigFile(cfgFile string) {
	c.configFile = cfgFile

}

func (c *Config) readInConfig() error {
	var (
		err error
	)
	c.data, err = os.ReadFile(c.configFile)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) Binding(out interface{}) error {
	if err := c.readInConfig(); err != nil {
		return err
	}
	switch c.configType {
	case yamlConfig:
		if err := yaml.Unmarshal(c.data, out); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported config type: %s", c.configType)
	}
	return nil
}

type DefaultOptions struct {
	Collector CollectorOptions `yaml:"collector"`
	Query     QueryOptions     `yaml:"query"`
}

type CollectorOptions struct {
	Listen    string `yaml:"listen"`
	MaxTraces int    `yaml:"max-traces" mapstructure:"max-traces"`
}

type QueryOptions struct {
	Listen string `yaml:"listen"`
}

// Valid 校验配置
func (c *Config) Valid() error {
	return nil
}

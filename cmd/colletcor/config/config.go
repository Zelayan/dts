package config

type StoreType = string

type Config struct {
	Default   DefaultOptions
	StoreType StoreType
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

package options

import (
	"github.com/Zelayan/dts/cmd/colletcor/config"
	"github.com/Zelayan/dts/pkg/collector"
	"github.com/Zelayan/dts/pkg/store"
)

const (
	defaultListen = 3001
	defaultConfig = "/etc/dts/config.yaml"
)

type Options struct {
	ComponentConfig config.Config
	// 存储接口
	Factory store.ShareDaoFactory

	//Collector
	Collector collector.CollectorInterface
	// 默认的配置文件
	ConfigFile string
}

func NewOptions() (*Options, error) {
	return &Options{}, nil
}

func (o *Options) Complete() error {
	// TODO register config
	if o.ComponentConfig.StoreType == "" {
		o.ComponentConfig.StoreType = store.DefaultStoreType
	}
	if o.ComponentConfig.Default.Listen == "" {
		o.ComponentConfig.Default.Listen = "0.0.0.0:3001"
	}

	if err := o.register(); err != nil {
		return err
	}

	o.Collector = collector.New(o.ComponentConfig, o.Factory)
	return nil
}

func (o *Options) register() error {
	// 注册存储
	if err := o.registerStore(); err != nil {
		return err
	}

	// 注册其他依赖
	return nil
}

func (o *Options) registerStore() error {
	factory, err := store.NewDaoFactory(o.ComponentConfig.StoreType)
	if err != nil {
		return err
	}
	o.Factory = factory
	return err
}

func (o *Options) Validate() error {
	// TODO
	return nil
}

package options

import (
	"github.com/Zelayan/dts/cmd/collector/config"
	"github.com/Zelayan/dts/pkg/collector"
	"github.com/Zelayan/dts/pkg/store"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	DefaultListen      = "0.0.0.0:3001"
	DefaultConfig      = "/etc/dts/config.yaml"
	DefaultQueryListen = "0.0.0.0:4001"
)

type Options struct {
	ComponentConfig config.Config
	// 存储接口
	Factory store.ShareDaoFactory

	//Collector
	Collector collector.CollectorInterface
	// 默认的配置文件
	ConfigFile string

	HttpEngine *gin.Engine
}

func NewOptions() (*Options, error) {
	return &Options{
		HttpEngine: gin.Default(),
	}, nil
}

func (o *Options) Complete() error {
	if len(o.ConfigFile) == 0 {
		if cfgFile := os.Getenv("ConfigFile"); cfgFile != "" {
			o.ConfigFile = cfgFile
		} else {
			o.ConfigFile = DefaultConfig
		}
	}
	c := config.New()
	c.SetConfigType("yaml")
	c.SetConfigFile(o.ConfigFile)
	if err := c.Binding(&o.ComponentConfig); err != nil {
		return err
	}
	if o.ComponentConfig.StoreType == "" {
		o.ComponentConfig.StoreType = store.DefaultStoreType
	}
	if o.ComponentConfig.Default.Collector.Listen == "" {
		o.ComponentConfig.Default.Collector.Listen = DefaultListen
	}
	if o.ComponentConfig.Default.Query.Listen == "" {
		o.ComponentConfig.Default.Query.Listen = DefaultQueryListen
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
	factory, err := store.NewDaoFactory(o.ComponentConfig)
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

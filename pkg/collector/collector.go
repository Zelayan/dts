package collector

import (
	"github.com/Zelayan/dts/cmd/colletcor/config"
	"github.com/Zelayan/dts/pkg/collector/span"
	"github.com/Zelayan/dts/pkg/store"
)

type CollectorInterface interface {
	span.SpanGetter
}

type collector struct {
	cc      config.Config
	factory store.ShareDaoFactory
}

func (c *collector) Span() span.Interface {
	return span.NewSpan(c.cc, c.factory)
}

func New(cfg config.Config, f store.ShareDaoFactory) CollectorInterface {
	return &collector{
		cc:      cfg,
		factory: f,
	}
}

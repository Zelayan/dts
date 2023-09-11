package query

import (
	"github.com/Zelayan/dts/cmd/colletcor/options"
	"github.com/Zelayan/dts/pkg/collector"
	"github.com/gin-gonic/gin"
)

type queryRouter struct {
	c collector.CollectorInterface
}

func NewRouter(o *options.Options) {
	s := &queryRouter{
		c: o.Collector,
	}
	s.initRouter(o.HttpEngine)
}

func (q *queryRouter) initRouter(httpEngine *gin.Engine) {
	queryGroup := httpEngine.Group("/query")
	{
		queryGroup.GET("/services", q.getServices)
	}
}

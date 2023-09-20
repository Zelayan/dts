package query

import (
	"github.com/Zelayan/dts/api/server/httputils"
	"github.com/gin-gonic/gin"
)

type StructuredResponse struct {
	Data   interface{}       `json:"data"`
	Total  int               `json:"total"`
	Limit  int               `json:"limit"`
	Offset int               `json:"offset"`
	Error  []structuredError `json:"error"`
}

type structuredError struct {
	Code    int    `json:"code,omitempty"`
	Msg     string `json:"msg"`
	TraceID string `json:"trace_id,omitempty"`
}

func (q *queryRouter) getServices(c *gin.Context) {
	r := httputils.NewResponse()
	var (
		err error
	)
	r.Result, err = q.c.Span().ListService(c.Request.Context())
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}
	httputils.SetSuccess(c, r)
}

func (q *queryRouter) getTraces(c *gin.Context) {
	r := httputils.NewResponse()
	var (
		err error
	)
	traceId := c.Param("traceId")
	r.Result, err = q.c.Span().GetTraces(c.Request.Context(), traceId)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

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

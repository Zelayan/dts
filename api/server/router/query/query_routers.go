package query

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	service, err := q.c.Span().ListService(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, service)
}

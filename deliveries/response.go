package deliveries

import (
	"github.com/gin-gonic/gin"
)

type ResponseStatus string

const (
	STATUS_SUCCESS ResponseStatus = "success"
	STATUS_ERROR   ResponseStatus = "error"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func responseSuccess(c *gin.Context, statusCode int, msg string, result interface{}) {
	r := new(Response)
	r.Status = string(STATUS_SUCCESS)
	r.Message = msg
	r.Result = result
	c.AbortWithStatusJSON(statusCode, *r)
}

func responseError(c *gin.Context, statusCode int, err error) {
	r := new(Response)
	r.Status = string(STATUS_ERROR)
	r.Message = err.Error()
	c.AbortWithStatusJSON(statusCode, *r)
}

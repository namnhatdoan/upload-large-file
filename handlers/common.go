package handlers

import (
	"github.com/gin-gonic/gin"
	"io"
	"uploadLargeFile/constants"
)

func bindingBodyData(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBind(req); err != nil {
		switch {
		case err == io.EOF:
			responseBadRequest(c, constants.RequestBodyEmpty)
			return false
		default:
			responseBadRequest(c, err.Error())
			return false
		}
	}
	return true
}

func bindingQueryData(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBindQuery(req); err != nil {
		responseBadRequest(c, err.Error())
		return false
	}
	return true
}

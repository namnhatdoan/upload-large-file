package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"uploadLargeFile/constants"
)

type Response struct {
	Code int8        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func responseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: constants.Success,
		Msg: constants.SuccessMessage,
		Data: data,
	})
}

func responseBadRequest(c *gin.Context, err string) {
	response := Response{
		Code: constants.BadRequest,
		Msg: constants.BadRequestMessage,
		Data: err,
	}
	c.JSON(http.StatusOK, &response)
}

func responseGeneralFailure(c *gin.Context, err string) {
	c.JSON(http.StatusOK, &Response{
		Code: constants.GeneralFailure,
		Msg: constants.GeneralFailureMessage,
		Data: err,
	})
}

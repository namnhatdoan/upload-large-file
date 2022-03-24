package handlers

import "github.com/gin-gonic/gin"

type Handler interface {
	UploadData(c *gin.Context)
	GetData(c *gin.Context)
}

package main

import (
	"fmt"
	"log"
	"time"
	"uploadLargeFile/handlers"
	"uploadLargeFile/services"
	"uploadLargeFile/settings"
)
import "github.com/gin-gonic/gin"
import "github.com/gin-contrib/cors"


func initRouter() *gin.Engine {
	router := gin.Default()
	initCORS(router)

	return router
}

func initCORS(r *gin.Engine) {
	config := cors.DefaultConfig()

	config.AllowCredentials = true
	config.AllowWebSockets = true
	config.AllowAllOrigins = true
	config.ExposeHeaders = []string{"Content-Disposition"}
	config.MaxAge = 12 * time.Hour

	r.Use(cors.New(config))
}

func initHandler() handlers.Handler {
	svc := new(services.OHLCServiceImpl)
	return handlers.InitHandlerImpl(svc)
}

func main() {
	router := initRouter()
	handler := initHandler()

	router.GET("/data", handler.GetData)
	router.POST("/data", handler.UploadData)

	if err := router.Run(fmt.Sprintf(":%v", settings.RestPort)); err != nil {
		log.Fatal(err)
	}
}

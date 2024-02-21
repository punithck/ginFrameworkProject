package main

import (
	"ginFrameworkProject/controller"
	"ginFrameworkProject/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	// the server includes both Logger and Recovery
	// Logger print it in logger file
	server := gin.Default()
	server.GET("/videos", func(ctx *gin.Context) {
		// Here we are just calling controller to return response
		ctx.JSON(200, videoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		// Here we are just calling controller to return response
		ctx.JSON(200, videoController.Save(ctx))
	})
	server.Run(":8081")
}

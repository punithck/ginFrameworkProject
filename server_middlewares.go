package main

import (
	"ginFrameworkProject/controller"
	"ginFrameworkProject/middlewares"
	"ginFrameworkProject/service"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	// creating a file
	f, _ := os.Create("gin.log")
	// We are loading the defaultWriter multiWriter
	// Why do you do this? You are giving your code to framework, It is the process they defined for thi
	// We have to follow what the framework says. Frameworks are opinionated. No flexibility but to follow the
	// process, flow the designed.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	setupLogOutput()
	// Here we are creating new server
	// Previously we were using gin.Default(). It was internally using Logger() and Recovery()
	server := gin.New()
	// We are adding Recovery middleware which was implemented by gene
	// We can also use gin.Logger() to logger middleware
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())
	// If you want to print more information about the request, you can use middlewares gindump.Dump()
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

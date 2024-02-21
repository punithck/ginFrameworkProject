package main

import (
	"ginFrameworkProject/controller"
	"ginFrameworkProject/middlewares"
	"ginFrameworkProject/service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

// Server side validation is mandatory to follow 3 key rules to develop secure application
// Never trust user input, Never trust user input and never trust user input
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

	server.Static("/css", "./template/css")
	server.LoadHTMLGlob("templates/*.html")
	// We can group the apis based on API and view
	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			// Here we are just calling controller to return response
			ctx.JSON(200, videoController.FindAll())
		})
		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			// Here we can use validation part provided by the gin or we can write our custom validation
			// Here we are using golang struct tags
			v, err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, v)
			}
		})
	}
	viewRoutes := server.Group("/views")
	{
		// internally it takes care of auth also
		viewRoutes.GET("/videos", videoController.ShowAll)
		// I do not know how to skip authorization. Need to check
	}
	// port you can get from environment variable os.GetEnv("PORT)
	server.Run(":8081")
}

// We can easily run this application in docker. We just need Dockerfile.

// JWT

// Industry standard
// Secure mechanism to transfer claims between 2 parties
// Claims are encoded as JSON object and digitally signed with secret key
// single token can be used for multiple backends services
// No session managment (100 % stateless). No db requred, no inmemory key
// You can use JWT libraries

//JSON web token
//Header
//Payload
//signature

// In the video, he is deviating from the topic. He started talking about JWT token and how to use this in the
// webapplication

//GORM: It is used. ORM libraray for golang. developer friendly to interact with database
// supports, CRED, SQL builder,plugins based on gorm callbasckl,official support for SQLLite, Mysql,postregress and mysql server
// GORM needs special attention on how to use it.

// Swagger can be used to test our rest api. swagger supports gin framework. It can be integrated.

// We can use Ginko and Gomega(assertion api) golang framework for testing apis

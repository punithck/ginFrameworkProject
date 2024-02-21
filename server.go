package main

import "github.com/gin-gonic/gin"

func main() {
	// the server includes both Logger and Recovery
	// Logger print it in logger file
	server := gin.Default()
	// gin.Context holds request, response,index, fullPath and other information
	// Here we are passing function for handling the GET request
	server.GET("/test", func(ctx *gin.Context) {
		// gin.H is of type map[string]any
		// any is of type of interface
		// I can pass any value to the as second variable
		ctx.JSON(200, gin.H{
			"message": "OK",
		})

	})
	// Here we run the server in the port 8081
	// Server runs in debug mode, You can see that in the log. You see these things in log
	// [log type] port listening to. GET/PUT/POST request, response of each request(200, 400, 404)
	server.Run(":8081")
}

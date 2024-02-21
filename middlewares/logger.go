package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// You are writting middlewares. All middlewares Returns HandlerFunc.
// You have to use gin.Logger functions to write your formats
// Why do you do this? Because you are giving your code to framework to run. Framework understands this way
// Frameworks are opiniontated. They defined the middlewares to be like this. Hence use this.
func Logger() gin.HandlerFunc {
	//gin.Logger() You can refer this
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - %s - %s - %s \n",
			params.ClientIP,
			params.Method,
			params.Latency,
			params.TimeStamp.Format(time.RFC822))
	})
}

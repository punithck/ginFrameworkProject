package middlewares

import "github.com/gin-gonic/gin"

// It is the framework.
// They asked us to write middleware for BasicAuth like this. Multiple options are given.
// Choose one of the option. I am choosing BasicAuth
func BasicAuth() gin.HandlerFunc {
	// Any doubt look at the function definition(document). It accepts username and passwords in the
	// form of map[string][string]
	return gin.BasicAuth(gin.Accounts{
		"pu": "ni",
	})
}

package middleware

import "github.com/gin-gonic/gin"

func Authenticate() gin.HandlerFunc{
	return func(c *gin.Context) {
		clientToekn := c.Request.Header.Get("token")
		
	}
}
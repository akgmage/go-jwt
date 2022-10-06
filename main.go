package main

import (
	routes "go-jwt/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	router.GET("/api-1", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"Access granted for api-1"})
	})
	router.GET("api-2", func(c *gin.Context){
		c.JSON(200, gin.H{"succcess": "Access granted for api-2"})
	})
	router.Run(":" + port)
}
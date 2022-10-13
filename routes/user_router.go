package routes

import (
	controller "go-jwt/controllers"
	"go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

/// attach a global middleware to the router. i.e. the middleware attached through Use() 
/// will be included in the handlers chain for every single request
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
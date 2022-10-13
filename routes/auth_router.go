package routes

import (
	controller "go-jwt/controllers"

	"github.com/gin-gonic/gin"
)

/// Routes with no middleware attached
func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}
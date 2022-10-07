package controllers

import (
	"context"
	"go-jwt/helpers"
	helper "go-jwt/helpers"
	"go-jwt/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/akgmage/go-jwt/database"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()
func HashPassword() 

func VerifyPassword()

func Signup()

func Login()

func GetUsers()

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}
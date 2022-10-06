package database

import (
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {log.Fatal("Error loading .env file")}
	MongoDb := os.Getenv("MONGODB_url")
}
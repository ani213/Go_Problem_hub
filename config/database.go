package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// ConnectDB initializes MongoDB connection
func ConnectDB() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get MongoDB URI from .env
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI not set in .env")
	}

	// Define MongoDB client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB using mongo.Connect()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}

	// Verify connection with Ping
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	// Select database
	DB = client.Database(os.Getenv("DB_NAME"))

	fmt.Println("✅ Connected to MongoDB")
}

// GetCollection returns a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	if DB == nil {
		log.Fatal("❌ MongoDB is not connected! Call ConnectDB() first.")
	}
	return DB.Collection(collectionName)
}

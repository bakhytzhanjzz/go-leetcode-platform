package main

import (
	"log"
	"os"

	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/database"
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/internal/grpcclient"
	natsclient "github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/internal/nats"
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	database.InitDB()

	// Initialize gRPC user client
	userClient := grpcclient.NewUserClient("localhost:50051")

	// Initialize NATS publisher
	publisher := mustInitNATSPublisher("nats://localhost:4222")

	// Setup Gin router with CORS
	router := setupRouter(userClient, publisher)

	// Run server
	port := getPort()
	log.Printf("Submission service running on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

// setupRouter configures all routes and middleware
func setupRouter(userClient *grpcclient.UserClient, publisher *natsclient.Publisher) *gin.Engine {
	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.RegisterSubmissionRoutes(r, database.DB, userClient, publisher)
	return r
}

// mustInitNATSPublisher initializes NATS or exits on failure
func mustInitNATSPublisher(url string) *natsclient.Publisher {
	publisher, err := natsclient.NewPublisher(url)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	return publisher
}

// getPort returns the server port or default
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8082"
	}
	return port
}

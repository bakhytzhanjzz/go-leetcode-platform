package main

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/database"
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/internal/grpcclient"
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/internal/nats"
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/routes"

	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	database.InitDB()

	userClient := grpcclient.NewUserClient("localhost:50051")

	// Setup NATS publisher
	publisher, err := natsclient.NewPublisher("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}

	routes.RegisterSubmissionRoutes(r, database.DB, userClient, publisher)

	r.Run(":8082")
}

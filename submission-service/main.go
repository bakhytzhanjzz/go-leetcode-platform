package main

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/database"
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/internal/grpcclient"
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.InitDB()

	// Initialize User gRPC client BEFORE registering routes
	userClient := grpcclient.NewUserClient("localhost:50051")

	// Pass userClient to routes for injection into handlers
	routes.RegisterSubmissionRoutes(r, database.DB, userClient)

	r.Run(":8082")
}

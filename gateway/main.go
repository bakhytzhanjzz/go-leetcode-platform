package main

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/gateway/internal/grpcclient"
	"github.com/bakhytzhanjzz/go-leetcode-platform/gateway/middleware"
	"github.com/bakhytzhanjzz/go-leetcode-platform/gateway/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	userClient := grpcclient.NewUserClient("localhost:50051")

	// Public routes
	routes.RegisterUserRoutes(r)

	// Protected routes
	authMiddleware := middleware.AuthMiddleware(userClient)

	api := r.Group("/api", authMiddleware)
	routes.RegisterProblemRoutes(api)
	routes.RegisterSubmissionRoutes(api)

	r.Run(":8080") // Gateway exposed at 8080
}

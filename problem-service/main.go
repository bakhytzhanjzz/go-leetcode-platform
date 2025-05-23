package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/database"
	natsclient "github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/internal/nats"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/internal/pkg"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/models"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/repository"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/routes"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/server"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type SubmissionCreatedEvent struct {
	SubmissionID uint `json:"submission_id"`
	UserID       uint `json:"user_id"`
	ProblemID    uint `json:"problem_id"`
}

func main() {
	// --- Init database and auto-migrate models
	db := database.InitDB()
	db.AutoMigrate(&models.Problem{})

	// --- Init Redis client
	redisClient := pkg.NewRedisClient("localhost:6379", "", 0)

	// --- Setup repositories
	problemRepo := repository.NewProblemRepo(db, redisClient)
	categoryRepo := repository.NewCategoryRepo(db)

	// --- Init Gin router with CORS middleware
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Change as needed
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// --- Register routes
	routes.RegisterProblemRoutes(router, problemRepo, categoryRepo)

	// --- Init NATS subscriber
	subscriber, err := natsclient.NewSubscriber("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}

	// --- Handle submission.created event (optional use-case)
	err = subscriber.Subscribe("submission.created", func(msg []byte) {
		var event SubmissionCreatedEvent
		if err := json.Unmarshal(msg, &event); err != nil {
			log.Printf("Failed to parse submission.created event: %v", err)
			return
		}
		log.Printf("Received submission.created event: %+v", event)
		fmt.Printf("Processing problem ID %d related to submission %d\n", event.ProblemID, event.SubmissionID)
	})
	if err != nil {
		log.Fatalf("Failed to subscribe to submission.created: %v", err)
	}

	// --- Handle submission.judged event (real logic)
	err = subscriber.HandleSubmissionJudged(problemRepo)
	if err != nil {
		log.Fatalf("Failed to subscribe to submission.judged: %v", err)
	}

	// --- Start servers (HTTP and gRPC)
	go func() {
		if err := router.Run(":8081"); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()
	server.StartGRPCServer(db)
}

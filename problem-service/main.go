package main

import (
	"encoding/json"
	"fmt"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/routes"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/server"
	"log"

	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/database"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/handlers"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/internal/nats"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/repository"
	"github.com/gin-gonic/gin"
)

type SubmissionCreatedEvent struct {
	SubmissionID uint `json:"submission_id"`
	UserID       uint `json:"user_id"`
	ProblemID    uint `json:"problem_id"`
}

func main() {
	db := database.InitDB()
	r := gin.Default()
	repo := repository.NewProblemRepo(db)
	handler := handlers.NewProblemHandler(repo)

	// Register routes with handler
	routes.RegisterProductRoutes(r, db)

	// Setup NATS subscriber
	subscriber, err := natsclient.NewSubscriber("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}

	err = subscriber.Subscribe("submission.created", func(msg []byte) {
		var event SubmissionCreatedEvent
		if err := json.Unmarshal(msg, &event); err != nil {
			log.Printf("Failed to parse submission.created event: %v", err)
			return
		}
		log.Printf("Received submission.created event: %+v", event)

		// Example action: update problem stats or print log
		// e.g., repo.IncrementSubmissionCount(event.ProblemID)

		fmt.Printf("You can process problem ID %d related to new submission %d\n", event.ProblemID, event.SubmissionID)
	})
	if err != nil {
		log.Fatalf("Failed to subscribe to submission.created: %v", err)
	}

	r.Run(":8081")
	server.StartGRPCServer(db)
}

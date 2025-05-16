package main

import (
	"encoding/json"
	"fmt"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/models"
	"log"

	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/database"
	natsclient "github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/internal/nats"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/repository"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/routes"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/server"
	"github.com/gin-gonic/gin"
)

type SubmissionCreatedEvent struct {
	SubmissionID uint `json:"submission_id"`
	UserID       uint `json:"user_id"`
	ProblemID    uint `json:"problem_id"`
}

func main() {
	db := database.InitDB()
	db.AutoMigrate(&models.Problem{})
	r := gin.Default()

	problemRepo := repository.NewProblemRepo(db)
	categoryRepo := repository.NewCategoryRepo(db)

	// Pass actual repos instead of raw db
	routes.RegisterProblemRoutes(r, problemRepo, categoryRepo)

	subscriber, err := natsclient.NewSubscriber("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}

	// Listen for submission.created (optional behavior)
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

	// Real use-case listener
	err = subscriber.HandleSubmissionJudged(problemRepo)
	if err != nil {
		log.Fatalf("Failed to subscribe to submission.judged: %v", err)
	}

	r.Run(":8081")
	server.StartGRPCServer(db)
}

package main

import (
	"fmt"

	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/internal/grpcclient"
)

func main() {
	userClient := grpcclient.NewUserClient("localhost:50051")

	// Replace with a valid JWT token
	testToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3NDc0MTk0NDB9.gaeKfENANsYd9nGvdBqs3-si8oWAiS430lxwqSC8qI4"

	userID, valid, err := userClient.ValidateToken(testToken)
	fmt.Println("Valid:", valid)
	fmt.Println("UserID:", userID)
	fmt.Println("Error:", err)
}

package main

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/database"
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.InitDB()
	routes.RegisterSubmissionRoutes(r, database.DB)
	r.Run(":8082")
}

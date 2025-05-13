package main

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/database"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitDB()
	r := gin.Default()
	routes.RegisterProductRoutes(r, db)

	r.Run(":8081")
}

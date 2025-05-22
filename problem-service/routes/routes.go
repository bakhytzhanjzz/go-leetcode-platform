package routes

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/handlers"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/repository"
	"github.com/gin-gonic/gin"
)

func RegisterProblemRoutes(r *gin.Engine, problemRepo *repository.ProblemRepo, categoryRepo *repository.CategoryRepo) {
	problemHandler := handlers.NewProblemHandler(problemRepo)
	categoryHandler := &handlers.CategoryHandler{Repo: categoryRepo}

	r.POST("/problems", problemHandler.Create)
	r.GET("/problems/:id", problemHandler.Get)
	r.PATCH("/problems/:id", problemHandler.Update)
	r.DELETE("/problems/:id", problemHandler.Delete)
	r.GET("/problems", problemHandler.List)

	r.POST("/categories", categoryHandler.Create)
	r.GET("/categories/:id", categoryHandler.Get)
	r.PATCH("/categories/:id", categoryHandler.Update)
	r.DELETE("/categories/:id", categoryHandler.Delete)
	r.GET("/categories", categoryHandler.List)
}

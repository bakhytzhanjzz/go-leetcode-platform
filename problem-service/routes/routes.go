package routes

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/handlers"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterProductRoutes(r *gin.Engine, db *gorm.DB) {
	problemRepo := &repository.ProblemRepo{DB: db}
	problemHandler := &handlers.ProblemHandler{Repo: problemRepo}

	r.POST("/problems", problemHandler.Create)
	r.GET("/problems/:id", problemHandler.Get)
	r.PATCH("/problems/:id", problemHandler.Update)
	r.DELETE("/problems/:id", problemHandler.Delete)
	r.GET("/problems", problemHandler.List)

	categoryRepo := &repository.CategoryRepo{DB: db}
	categoryHandler := &handlers.CategoryHandler{Repo: categoryRepo}

	r.POST("/categories", categoryHandler.Create)
	r.GET("/categories/:id", categoryHandler.Get)
	r.PATCH("/categories/:id", categoryHandler.Update)
	r.DELETE("/categories/:id", categoryHandler.Delete)
	r.GET("/categories", categoryHandler.List)
}

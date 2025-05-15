package routes

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/handlers"
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterSubmissionRoutes(r *gin.Engine, db *gorm.DB) {
	repo := &repository.SubmissionRepo{DB: db}
	handler := &handlers.SubmissionHandler{Repo: repo}

	r.POST("/submissions", handler.Create)
	r.GET("/submissions/:id", handler.Get)
	r.GET("/submissions", handler.List)
	r.PATCH("/submissions/:id", handler.UpdateStatus)

}

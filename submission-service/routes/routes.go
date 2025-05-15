package routes

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/handlers"
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/internal/grpcclient"
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterSubmissionRoutes(r *gin.Engine, db *gorm.DB, userClient *grpcclient.UserClient) {
	submissionRepo := repository.NewSubmissionRepo(db)
	submissionHandler := handlers.SubmissionHandler{
		Repo:       submissionRepo,
		UserClient: userClient,
	}

	r.POST("/submissions", submissionHandler.Create)
	r.GET("/submissions/:id", submissionHandler.Get)
	r.GET("/submissions", submissionHandler.List)
	r.PATCH("/submissions/:id", submissionHandler.UpdateStatus)
}

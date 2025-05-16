package handlers

import (
	"fmt"
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/internal/grpcclient"
	natsclient "github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/internal/nats"
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/models"
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/repository"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type SubmissionHandler struct {
	Repo       *repository.SubmissionRepo
	UserClient *grpcclient.UserClient
	Publisher  *natsclient.Publisher
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (h *SubmissionHandler) Create(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	userID, valid, errMsg := h.UserClient.ValidateToken(token)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + errMsg})
		return
	}

	var submission models.Submission
	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	submission.ID = 0
	submission.UserID = uint(userID)
	submission.Status = "Pending"

	if err := h.Repo.Create(&submission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create submission"})
		return
	}

	// Publish event: submission.created
	eventPayload := fmt.Sprintf(`{"submission_id": %d, "user_id": %d, "problem_id": %d}`, submission.ID, submission.UserID, submission.ProblemID)
	h.Publisher.Publish("submission.created", []byte(eventPayload))

	// Simulate evaluation asynchronously
	go func(sub models.Submission) {
		time.Sleep(3 * time.Second)
		statuses := []string{"Accepted", "Wrong Answer", "Runtime Error"}
		finalStatus := statuses[rand.Intn(len(statuses))]
		sub.Status = finalStatus
		_ = h.Repo.Update(&sub)

		// Publish event: submission.judged
		judgedPayload := fmt.Sprintf(`{"submission_id": %d, "user_id": %d, "problem_id": %d, "status": "%s"}`, sub.ID, sub.UserID, sub.ProblemID, finalStatus)
		h.Publisher.Publish("submission.judged", []byte(judgedPayload))

		// Publish event: user.notify
		notifyMsg := fmt.Sprintf(`{"user_id": %d, "message": "Your submission %d was judged: %s"}`, sub.UserID, sub.ID, finalStatus)
		h.Publisher.Publish("user.notify", []byte(notifyMsg))
	}(submission)

	c.JSON(http.StatusCreated, submission)
}

func (h *SubmissionHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	submission, err := h.Repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Submission not found"})
		return
	}
	c.JSON(http.StatusOK, submission)
}

func (h *SubmissionHandler) List(c *gin.Context) {
	submissions, err := h.Repo.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not list submissions"})
		return
	}
	c.JSON(http.StatusOK, submissions)
}

func (h *SubmissionHandler) UpdateStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var update struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	submission, err := h.Repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Submission not found"})
		return
	}

	submission.Status = update.Status
	if err := h.Repo.Update(submission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update status"})
		return
	}

	c.JSON(http.StatusOK, submission)
}

func NewSubmissionHandler(repo *repository.SubmissionRepo, userClient *grpcclient.UserClient, publisher *natsclient.Publisher) *SubmissionHandler {
	return &SubmissionHandler{
		Repo:       repo,
		UserClient: userClient,
		Publisher:  publisher,
	}
}

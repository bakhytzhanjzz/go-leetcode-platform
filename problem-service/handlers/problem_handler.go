package handlers

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/internal/pkg"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/models"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProblemHandler struct {
	Repo  *repository.ProblemRepo
	Cache *pkg.RedisClient
}

func (h *ProblemHandler) Create(c *gin.Context) {
	var problem models.Problem
	if err := c.ShouldBindJSON(&problem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure that category_id is valid (greater than 0)
	if problem.CategoryID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category ID is required and must be greater than 0"})
		return
	}

	if err := h.Repo.Create(&problem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create problem"})
		return
	}

	c.JSON(http.StatusCreated, problem)
}

func (h *ProblemHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	problem, err := h.Repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Problem not found"})
		return
	}
	c.JSON(http.StatusOK, problem)
}

func (h *ProblemHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var problem models.Problem
	if err := c.ShouldBindJSON(&problem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	problem.ID = uint(id)
	if err := h.Repo.Update(&problem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update problem"})
		return
	}
	c.JSON(http.StatusOK, problem)
}

func (h *ProblemHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete problem"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *ProblemHandler) List(c *gin.Context) {
	problems, err := h.Repo.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not list problems"})
		return
	}
	c.JSON(http.StatusOK, problems)
}
func NewProblemHandler(repo *repository.ProblemRepo) *ProblemHandler {
	return &ProblemHandler{
		Repo:  repo,
		Cache: repo.Cache, // if Cache is a field on ProblemRepo
	}
}

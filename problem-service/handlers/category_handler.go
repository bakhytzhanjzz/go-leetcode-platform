package handlers

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/models"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	Repo *repository.CategoryRepo
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var cat models.Category
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Repo.Create(&cat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create category"})
		return
	}
	c.JSON(http.StatusCreated, cat)
}

func (h *CategoryHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	cat, err := h.Repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, cat)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var cat models.Category
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cat.ID = uint(id)
	if err := h.Repo.Update(&cat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update category"})
		return
	}
	c.JSON(http.StatusOK, cat)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete category"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *CategoryHandler) List(c *gin.Context) {
	cats, err := h.Repo.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not list categories"})
		return
	}
	c.JSON(http.StatusOK, cats)
}

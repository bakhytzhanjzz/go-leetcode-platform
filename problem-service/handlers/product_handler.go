package handlers

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/models"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	Repo *repository.ProductRepo
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Repo.Create(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	c.JSON(http.StatusCreated, p)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p, err := h.Repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	p.ID = uint(id)
	if err := h.Repo.Update(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.Repo.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not list products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

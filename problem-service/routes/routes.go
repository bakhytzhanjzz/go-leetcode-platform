package routes

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/handlers"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterProductRoutes(r *gin.Engine, db *gorm.DB) {
	repo := &repository.ProductRepo{DB: db}
	handler := &handlers.ProductHandler{Repo: repo}

	r.POST("/products", handler.CreateProduct)
	r.GET("/products/:id", handler.GetProduct)
	r.PATCH("/products/:id", handler.UpdateProduct)
	r.DELETE("/products/:id", handler.DeleteProduct)
	r.GET("/products", handler.ListProducts)
}

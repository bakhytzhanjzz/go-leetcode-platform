package repository

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/models"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	DB *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{DB: db}
}

func (r *CategoryRepo) Create(category *models.Category) error {
	return r.DB.Create(category).Error
}

func (r *CategoryRepo) GetByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.DB.First(&category, id).Error
	return &category, err
}

func (r *CategoryRepo) Update(category *models.Category) error {
	return r.DB.Save(category).Error
}

func (r *CategoryRepo) Delete(id uint) error {
	return r.DB.Delete(&models.Category{}, id).Error
}

func (r *CategoryRepo) List() ([]models.Category, error) {
	var categories []models.Category
	err := r.DB.Find(&categories).Error
	return categories, err
}

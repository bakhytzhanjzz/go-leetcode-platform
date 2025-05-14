package repository

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/models"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	DB *gorm.DB
}

func (r *CategoryRepo) Create(cat *models.Category) error {
	return r.DB.Create(cat).Error
}

func (r *CategoryRepo) GetByID(id uint) (*models.Category, error) {
	var cat models.Category
	err := r.DB.First(&cat, id).Error
	return &cat, err
}

func (r *CategoryRepo) Update(cat *models.Category) error {
	return r.DB.Save(cat).Error
}

func (r *CategoryRepo) Delete(id uint) error {
	return r.DB.Delete(&models.Category{}, id).Error
}

func (r *CategoryRepo) List() ([]models.Category, error) {
	var cats []models.Category
	err := r.DB.Find(&cats).Error
	return cats, err
}

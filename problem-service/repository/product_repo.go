package repository

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/models"
	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

func (repo *ProductRepo) Create(product *models.Product) error {
	return repo.DB.Create(product).Error
}

func (repo *ProductRepo) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	err := repo.DB.First(&product, id).Error
	return &product, err
}

func (repo *ProductRepo) Update(product *models.Product) error {
	return repo.DB.Save(product).Error
}

func (repo *ProductRepo) Delete(id uint) error {
	return repo.DB.Delete(&models.Product{}, id).Error
}

func (repo *ProductRepo) List() ([]models.Product, error) {
	var products []models.Product
	err := repo.DB.Find(&products).Error
	return products, err
}

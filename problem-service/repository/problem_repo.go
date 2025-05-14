package repository

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/models"
	"gorm.io/gorm"
)

type ProblemRepo struct {
	DB *gorm.DB
}

func (r *ProblemRepo) Create(problem *models.Problem) error {
	return r.DB.Create(problem).Error
}

func (r *ProblemRepo) GetByID(id uint) (*models.Problem, error) {
	var problem models.Problem
	err := r.DB.Preload("Category").First(&problem, id).Error
	return &problem, err
}

func (r *ProblemRepo) Update(problem *models.Problem) error {
	return r.DB.Save(problem).Error
}

func (r *ProblemRepo) Delete(id uint) error {
	return r.DB.Delete(&models.Problem{}, id).Error
}

func (r *ProblemRepo) List() ([]models.Problem, error) {
	var problems []models.Problem
	err := r.DB.Preload("Category").Find(&problems).Error
	return problems, err
}

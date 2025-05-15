package repository

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/models"
	"gorm.io/gorm"
)

type SubmissionRepo struct {
	DB *gorm.DB
}

func NewSubmissionRepo(db *gorm.DB) *SubmissionRepo {
	return &SubmissionRepo{DB: db}
}

func (r *SubmissionRepo) Create(submission *models.Submission) error {
	return r.DB.Create(submission).Error
}

func (r *SubmissionRepo) GetByID(id uint) (*models.Submission, error) {
	var submission models.Submission
	err := r.DB.First(&submission, id).Error
	return &submission, err
}

func (r *SubmissionRepo) List() ([]models.Submission, error) {
	var submissions []models.Submission
	err := r.DB.Find(&submissions).Error
	return submissions, err
}

func (r *SubmissionRepo) Update(sub *models.Submission) error {
	return r.DB.Save(sub).Error
}

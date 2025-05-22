package repository

import (
	"encoding/json"
	"fmt"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/internal/pkg"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/models"
	"gorm.io/gorm"
	"time"
)

type ProblemRepo struct {
	DB    *gorm.DB
	Cache *pkg.RedisClient
}

func NewProblemRepo(db *gorm.DB, cache *pkg.RedisClient) *ProblemRepo {
	return &ProblemRepo{DB: db, Cache: cache}
}

func (r *ProblemRepo) GetByID(id uint) (*models.Problem, error) {
	cacheKey := fmt.Sprintf("problem:%d", id)

	// Try cache
	cached, err := r.Cache.Get(cacheKey)
	if err == nil {
		var p models.Problem
		if err := json.Unmarshal([]byte(cached), &p); err == nil {
			return &p, nil
		}
	}

	// Cache miss → DB
	var problem models.Problem
	err = r.DB.Preload("Category").First(&problem, id).Error
	if err != nil {
		return nil, err
	}

	// Store in Redis
	data, _ := json.Marshal(problem)
	_ = r.Cache.Set(cacheKey, string(data), 5*time.Minute)

	return &problem, nil
}

func (r *ProblemRepo) Create(problem *models.Problem) error {
	err := r.DB.Create(problem).Error
	if err != nil {
		return err
	}
	// Optionally cache the new problem
	cacheKey := fmt.Sprintf("problem:%d", problem.ID)
	data, _ := json.Marshal(problem)
	_ = r.Cache.Set(cacheKey, string(data), 5*time.Minute)
	return nil
}

func (r *ProblemRepo) Update(problem *models.Problem) error {
	err := r.DB.Save(problem).Error
	if err != nil {
		return err
	}
	// Update cache
	cacheKey := fmt.Sprintf("problem:%d", problem.ID)
	data, _ := json.Marshal(problem)
	_ = r.Cache.Set(cacheKey, string(data), 5*time.Minute)
	return nil
}

func (r *ProblemRepo) Delete(id uint) error {
	err := r.DB.Delete(&models.Problem{}, id).Error
	if err != nil {
		return err
	}
	// Invalidate cache
	cacheKey := fmt.Sprintf("problem:%d", id)
	_ = r.Cache.Del(cacheKey)
	return nil
}

func (r *ProblemRepo) List() ([]models.Problem, error) {
	var problems []models.Problem
	err := r.DB.Preload("Category").Find(&problems).Error
	return problems, err
}

func (r *ProblemRepo) IncrementSubmissionCount(problemID uint) error {
	return r.DB.Model(&models.Problem{}).
		Where("id = ?", problemID).
		UpdateColumn("submission_count", gorm.Expr("submission_count + ?", 1)).
		Error
}

func (r *ProblemRepo) IncrementAcceptedCount(problemID uint) error {
	return r.DB.Model(&models.Problem{}).
		Where("id = ?", problemID).
		UpdateColumn("accepted_count", gorm.Expr("accepted_count + ?", 1)).
		Error
}

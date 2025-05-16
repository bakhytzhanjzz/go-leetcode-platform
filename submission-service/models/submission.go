package models

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	ProblemID uint   `json:"problem_id"`
	Code      string `json:"code"`
	Language  string `json:"language"`
	Status    string `json:"status"`
}

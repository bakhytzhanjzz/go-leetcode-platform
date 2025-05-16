package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string    `json:"name"`
	Problems []Problem `gorm:"foreignKey:CategoryID"`
}

type Problem struct {
	gorm.Model
	Title           string `json:"title"`
	Description     string `json:"description"`
	SubmissionCount uint
	AcceptedCount   uint
	Difficulty      string   `json:"difficulty"`
	CategoryID      uint     `json:"category_id"`
	Category        Category `gorm:"foreignKey:CategoryID"`
}

package entities

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	ID        uint           `json:"id" gorm:"column:id;primaryKey"`
	Title     string         `json:"title" gorm:"column:title"`
	Content   string         `json:"content" gorm:"column:content"`
	Author    string         `json:"author" gorm:"column:author"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

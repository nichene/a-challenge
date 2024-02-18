package postgres

import (
	"gorm.io/gorm"
)

type ParentsRepository struct {
	db *gorm.DB
}

func NewParentsRepository(db *gorm.DB) *ParentsRepository {
	return &ParentsRepository{db}
}

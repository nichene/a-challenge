package postgres

import (
	"gorm.io/gorm"
)

type PersonsRepository struct {
	db *gorm.DB
}

func NewPersonsRepository(db *gorm.DB) *PersonsRepository {
	return &PersonsRepository{db}
}

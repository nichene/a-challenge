package postgres

import (
	"context"
	"stone_challenge/internal/person"
	"time"

	"gorm.io/gorm"
)

type ParentsRepository struct {
	db *gorm.DB
}

func NewParentsRepository(db *gorm.DB) *ParentsRepository {
	return &ParentsRepository{db}
}

func (r *ParentsRepository) Create(ctx context.Context, personID, parentID int) (person.Parent, error) {
	parent := &person.Parent{
		PersonID:  personID,
		ParentID:  parentID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err := r.db.Create(parent).Error
	if err != nil {
		return person.Parent{}, err
	}

	foundParent := &person.Parent{}
	err = r.db.Model(&person.Parent{}).Where("person_id = ?", personID).Where("parent_id = ?", parentID).First(foundParent).Error

	return *foundParent, err
}

func (r *ParentsRepository) Find(ctx context.Context, personID, parentID int) (person.Parent, error) {
	foundParent := &person.Parent{}
	result := r.db.Model(&person.Parent{}).Where("person_id = ?", personID).Where("parent_id = ?", parentID).First(foundParent).Error

	return *foundParent, result
}

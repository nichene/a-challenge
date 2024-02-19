package postgres

import (
	"context"
	"stone_challenge/internal/person"
	"time"

	"gorm.io/gorm"
)

type PersonsRepository struct {
	db *gorm.DB
}

func NewPersonsRepository(db *gorm.DB) *PersonsRepository {
	return &PersonsRepository{db}
}

func (r *PersonsRepository) FindByName(ctx context.Context, name string) (*person.Person, error) {
	foundPerson := &person.Person{}

	result := r.db.Model(&person.Person{}).Where("name = ?", name).First(foundPerson)

	return foundPerson, result.Error
}

func (r *PersonsRepository) Create(ctx context.Context, intent person.CreatePersonIntent) (*person.Person, error) {
	personToCreate := &person.Person{
		Name:      intent.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err := r.db.Create(personToCreate).Error
	if err != nil {
		return &person.Person{}, err
	}

	foundPerson := &person.Person{}
	err = r.db.Model(&person.Person{}).Where("name = ?", intent.Name).First(foundPerson).Error

	return foundPerson, err
}

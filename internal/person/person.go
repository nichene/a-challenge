package person

import (
	"context"

	"gorm.io/gorm"
)

func (s *Service) Create(ctx context.Context, intent CreatePersonIntent) (Person, error) {
	person, err := s.personsRepository.Create(ctx, intent)
	if err != nil {
		return Person{}, err
	}

	for _, parentName := range intent.Parents {
		parent, err := s.CreateParent(ctx, *person, parentName)
		if err == nil && parent.Name != "" {
			person.Parents = append(person.Parents, parent)
		}
	}

	for _, childName := range intent.Children {
		child, err := s.CreateChild(ctx, *person, childName)
		if err == nil && child.Name != "" {
			person.Children = append(person.Children, child)
		}
	}

	return *person, err
}

func (s *Service) CreateParent(ctx context.Context, person Person, parentName string) (Person, error) {
	parent, err := s.personsRepository.FindByName(ctx, parentName)
	if err == gorm.ErrRecordNotFound {
		parent, err = s.personsRepository.Create(ctx, CreatePersonIntent{Name: parentName})
	}

	if err != nil {
		return Person{}, err
	}

	_, err = s.parentsRepository.Find(ctx, person.ID, parent.ID)
	if err == gorm.ErrRecordNotFound {
		_, err = s.parentsRepository.Create(ctx, person.ID, parent.ID)
	}

	return *parent, err
}

func (s *Service) CreateChild(ctx context.Context, person Person, childName string) (Person, error) {
	child, err := s.personsRepository.FindByName(ctx, childName)
	if err == gorm.ErrRecordNotFound {
		child, err = s.personsRepository.Create(ctx, CreatePersonIntent{Name: childName})
	}

	if err != nil {
		return Person{}, err
	}

	_, err = s.parentsRepository.Find(ctx, child.ID, person.ID)
	if err == gorm.ErrRecordNotFound {
		_, err = s.parentsRepository.Create(ctx, child.ID, person.ID)
	}

	return *child, err
}

func (s *Service) FindByName(ctx context.Context, name string) (Person, error) {
	person, err := s.personsRepository.FindByName(ctx, name)

	return *person, err
}

// TO DO
func (s *Service) FindPersonAndParents(ctx context.Context, name string) (Person, error) {
	return Person{}, nil
}

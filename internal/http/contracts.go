package http

import (
	"stone_challenge/internal/health"
	"stone_challenge/internal/person"
)

type Service struct {
	Health        health.Checks
	PersonService person.Service
}

func NewService(health health.Checks, personService person.Service) *Service {
	return &Service{
		Health:        health,
		PersonService: personService,
	}
}

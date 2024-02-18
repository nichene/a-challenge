package handlers

import (
	"net/http"
	"stone_challenge/internal/person"

	"github.com/gofiber/fiber/v2"
)

type PersonHandler interface {
	Find(c *fiber.Ctx) error
}

type personHandler struct {
	personService person.Service
}

func NewPersonHandler(personService person.Service) PersonHandler {
	return &personHandler{
		personService: personService,
	}
}

func (h *personHandler) Find(ctx *fiber.Ctx) error {
	person, err := h.personService.Find(ctx.Context(), "name")
	if err != nil {

		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).SendString(person.Name)
}

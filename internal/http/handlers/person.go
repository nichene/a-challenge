package handlers

import (
	"log"
	"net/http"
	"stone_challenge/internal/person"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

type PersonHandler interface {
	Create(ctx *fiber.Ctx) error
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

func (h *personHandler) Create(ctx *fiber.Ctx) error {
	params := &person.CreatePersonIntent{}
	err := ctx.BodyParser(params)
	if err != nil {
		log.Default().Print("API - Unable to parse request body", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON("invalid params")
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		log.Default().Print("API - Invalid create person params")
		return ctx.Status(http.StatusBadRequest).JSON("invalid params")
	}

	_, err = h.personService.FindByName(ctx.Context(), params.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Default().Print("API - Error on finding person")
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	person, err := h.personService.Create(ctx.Context(), *params)
	if err != nil {
		log.Default().Print("API - Error creating person")
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(person)
}

func (h *personHandler) Find(ctx *fiber.Ctx) error {
	var name string
	err := ctx.QueryParser(&name)
	if err != nil {
		log.Default().Print("API - Unable to parse query param", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON("invalid query params")
	}

	person, err := h.personService.FindPersonAndParents(ctx.Context(), name)
	if err != nil {
		log.Default().Print("API - Error finding person")
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(person)
}

package routes

import (
	"stone_challenge/internal/http/handlers"

	"github.com/gofiber/fiber/v2"
)

// PersonRoutes contains persons endpoints.
func PersonRoutes(route *fiber.App, handler handlers.PersonHandler) {
	route.Post("/person", handler.Create)
	route.Get("/person", handler.Find)
}

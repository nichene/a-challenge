package http

import (
	"stone_challenge/internal/http/handlers"
	"stone_challenge/internal/http/routes"
)

func (s *Server) router() {
	routes.HealthRoute(s.httpServer, handlers.NewHealthHandler(s.service.Health))

	routes.PersonRoutes(s.httpServer, handlers.NewPersonHandler(s.service.PersonService))
}

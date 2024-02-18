package main

import (
	"context"
	"log"
	"os/signal"
	"stone_challenge/config"
	"stone_challenge/internal/health"
	"stone_challenge/internal/http"
	"stone_challenge/internal/person"
	postgresRepo "stone_challenge/internal/person/postgres"
	postgres "stone_challenge/pkg"
	"syscall"
)

func main() {
	cfg := config.LoadEnvVars()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	db, err := postgres.InitDatabase(ctx, cfg)

	httpService := http.NewService(
		health.NewHealthCheckService(),
		*person.NewService(
			postgresRepo.NewPersonsRepository(db),
			postgresRepo.NewParentsRepository(db),
		),
	)

	srv, err := http.NewServer(cfg.Port, http.WithService(httpService))
	if err != nil {
		log.Fatal("API - Server down", err)
	}

	log.Default().Print("API - New server up on port ", cfg.Port)

	<-ctx.Done()

	stop()
	log.Default().Print("API - Shutting down gracefully")

	err = srv.Close()
	if err != nil {
		log.Fatal("API - Shutdown error", err)
	}

	log.Default().Print("API - exiting")
}

package main

import (
	"context"
	"log"
	"os/signal"
	"stone_challenge/internal/health"
	"stone_challenge/internal/http"
	"stone_challenge/internal/person"
	"syscall"
)

func main() {
	envVars := loadEnvVars()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	httpService := http.NewService(health.NewHealthCheckService(), *person.NewService())

	srv, err := http.NewServer(envVars.Port, http.WithService(httpService))
	if err != nil {
		log.Fatal("API - Server down", err)
	}

	log.Default().Print("API - New server up on port ", envVars.Port)

	<-ctx.Done()

	stop()
	log.Default().Print("API - Shutting down gracefully")

	err = srv.Close()
	if err != nil {
		log.Fatal("API - Shutdown error", err)
	}

	log.Default().Print("API - exiting")
}

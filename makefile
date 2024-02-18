include .env

api:
	go run ./cmd/*.go

tests:
	go test ./...

db:
	docker-compose up -d postgres

GOOSE := goose -dir ./migrations postgres "host=$(DB_HOST) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable"
migrate-up:
	$(GOOSE) up

build-image:
	docker build -t stone-challeng:latest .
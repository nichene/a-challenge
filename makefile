
api:
	go run ./cmd/*.go

tests:
	go test ./...

build-image:
	docker build -t stone-challeng:latest .
.PHONY: run-dev debug build build-and-run build-dev-docker run-dev-docker

run-dev:
	go run cmd/main/main.go -APP_ENV=development -debug=false

debug:
	go run cmd/main/main.go -APP_ENV=development -debug

build:
	go build -o bin/spendaro-api cmd/main/main.go

build-and-run:
	go build -o bin/spendaro-api cmd/main/main.go
	./bin/spendaro-api -APP_ENV=development -debug=false

build-dev-docker:
	docker build -t spendaro-api -f ./docker/Dockerfile.dev .

run-dev-docker:
	docker run --name spendaro-api -p 8080:8080 spendaro-api

# Docker Compose
dev-compose:
	cd ./docker && docker-compose -f docker-compose-dev.yaml up

build-dev-compose:
	cd ./docker && docker-compose -f docker-compose-dev.yaml up --build

# Tests
test-verbose:
	go test -v ./...

test:
	go test ./...
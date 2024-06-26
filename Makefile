.PHONY: run-dev debug build build-and-run build-dev-docker run-dev-docker

run-dev:
	go run cmd/main/main.go -APP_ENV=development -DEBUG=false

debug:
	go run cmd/main/main.go -APP_ENV=development -DEBUG

build:
	go build -o bin/spendaro-api cmd/main/main.go

build-and-run:
	go build -o bin/spendaro-api cmd/main/main.go
	./bin/spendaro-api -APP_ENV=development -DEBUG=false

build-dev-docker:
	docker build -t spendaro-api -f ./docker/Dockerfile.dev .

run-dev-docker:
	docker run --name spendaro-api -p 8080:8080 spendaro-api

# Docker Compose
dev-compose:
	cd ./docker && docker-compose -p spendaro -f docker-compose-dev.yaml up -d

build-dev-compose:
	cd ./docker && docker-compose -p spendaro -f docker-compose-dev.yaml up --build -d

# Tests
test-verbose:
	go test -v ./...

test:
	go test ./...
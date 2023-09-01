.PHONY: build docker lint test

VERSION=$(shell git rev-parse --short HEAD)

build:
	go build -v ./...

docker:
	docker build -t service:$(VERSION) .

lint:
	golangci-lint run --fix

test:
	docker compose up -d --wait db
	migrate -source file://migrations -database 'postgres://service:password@localhost:5432/book?sslmode=disable' down -all
	go test ./test/... -v -count=1

.PHONY: build gen docker lint test

VERSION=$(shell git rev-parse --short HEAD)

build:
	go build -v ./...

gen:
	go run ./cmd/gen/main.go
	swag init -d api,internal/request,internal/response,internal/database/entities -o ./api/docs -g ./http.go

docker:
	docker build -t service:$(VERSION) .

lint:
	golangci-lint run --fix

test:
	docker compose up -d --wait db
	migrate -source file://migrations -database 'postgres://service:password@localhost:5432/book?sslmode=disable' down -all
	go test ./test/... -v -count=1

# Boilerplate

## Code structure

- Project layout

  - [golang-standards](https://github.com/golang-standards/project-layout)

- Storage client:

  - [gorm](https://gorm.io/docs/) => ORM library for Golang
  - [gen](https://gorm.io/gen/) => Type-safe DAO API without interface{}
  - [rueidis](https://github.com/redis/rueidis) => fast redis client

- Development tools:

  - [air](https://github.com/cosmtrek/air) => live reload for Go apps
  - [golangci-lint](https://golangci-lint.run/) => powerful linter
  - [golang-migrate](https://github.com/golang-migrate/migrate)
    => migration tool for go and cli

- Tech:

  - [echo](https://echo.labstack.com/) => web framework

- Utils:
  - [gcloud secret manager](https://cloud.google.com/secret-manager/docs/reference/libraries#client-libraries-install-go)
  - [zerolog](https://github.com/rs/zerolog) => zero allocation logger
  - [validator](https://github.com/go-playground/validator) => validator

### How to develop

#### Tool installation

```bash
$ go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
$ go install github.com/cosmtrek/air@latest
$ go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
$ go install github.com/swaggo/swag/cmd/swag@latest
```

#### Basic

```bash
# make sure you have dependencies up
$ docker compose up -d

# manual run
$ go run cmd/app/main.go

# hot reloading
$ air
```

#### Swagger

Please note that it's only for local environment

Served at: `http://localhost:8080/swagger/index.html`

```bash
$ swag init -d api,internal/request,internal/response,internal/database/entities -o ./api/docs -g ./http.go
```

#### Authentication

In local development, we send email in 'authorization' header

In development, uat and production environment, we use jwt auth

```bash
# local
$ curl localhost:8080/... -H 'authorization: admin@example.com'

# development, uat, production
$ curl localhost:8080/... -H 'authorization: Bearer ...'
```

#### Linter

❗️Please do this before raise the PRs ❗️

```bash
$ make lint
```

#### DAO Interface generation

❗️Please do this when you have updated the `entities` ❗️

```bash
$ go run cmd/gen/main.go
```

#### Migration

New migration, find out more on best practices [here](https://github.com/golang-migrate/migrate/blob/master/MIGRATIONS.md)

```bash
$ migrate create -digits 6 -dir migrations -ext sql -seq 'MIGRATION_NAME'
```

For testing only, in the code we have auto-migration in place.

```bash
$ migrate -source file://migrations -database 'postgres://service:password@localhost:5432/book?sslmode=disable' up
$ migrate -source file://migrations -database 'postgres://service:password@localhost:5432/book?sslmode=disable' down -all
$ migrate -source file://migrations -database 'postgres://service:password@localhost:5432/book?sslmode=disable' drop -f
```

#### Test

```bash
# make sure you have dependencies up
$ docker compose up -d

$ make test
```

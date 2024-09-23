# My Info Backend

Written in `Go` for Reduces Latency.

Things Supported:
 - CRUD Operation
 - JWT Based Authentication
 - JSON Validation
 - Paginated Response
 - SQLC For modules generation
 - Goose for Database Migration
 - Air Integration to Test in Dev ENV

### Setup

```bash
go install github.com/air-verse/air@latest
go get github.com/go-chi/jwtauth/v5
go get github.com/go-chi/chi/v5
go get golang.org/x/crypto/bcrypt
go get github.com/joho/godotenv
```

## Update Database

```bash
cd sql/schema
goose postgres postgres://<user_id>:<user_password>@localhost:5432/go_my_info up
```

## Generate Internal DB

```bash
sqlc generate
```

## Start Application

```bash
go build ./cmd/main.go && ./main
```

## Start Dev Environment

```bash
air
```

## Generate Swagger Documentation

```bash
swag init -o './api' -g './cmd/main.go'
```
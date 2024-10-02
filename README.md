# My Info Backend

<img align="right" width="180px" src="https://github.com/bytesByHarsh/go-my-info/blob/master/docs/assets/logo.png?raw=true">

Written in `Go` for Reduces Latency.

Things Supported:
 - CRUD Operation
 - JWT Based Authentication
 - JSON Validation
 - Paginated Response
 - SQLC For modules generation
 - Goose for Database Migration
 - Air Integration to Test in Dev ENV
 - Swagger Integration

## Swagger (Limited)
[http://localhost:8000/swagger/index.html](http://localhost:8000/swagger/index.html)

![Swagger Documentation](/docs/assets/swagger_1.png)

![Swagger Documentation](/docs/assets/swagger_2.png)

![Swagger Documentation of Models](/docs/assets/swagger_3.png)

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
go install github.com/swaggo/swag/cmd/swag@latest
swag init -o './api' -g './cmd/main.go' --parseDependency
swag fmt
```
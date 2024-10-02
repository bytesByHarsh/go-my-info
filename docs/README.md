# Go My Information Backend API

<img align="right" width="180px" src="https://github.com/bytesByHarsh/go-my-info/blob/master/docs/assets/logo.png?raw=true">

Simple Backend Application Written in `Go` to user personal account and cards details.

Things Supported:
 - CRUD Operation
 - JWT Based Authentication
 - JSON Validation
 - Paginated Response
 - SQLC For modules generation
 - Goose for Database Migration
 - Air Integration to Test in Dev ENV
 - Swagger Integration


![Swagger Documentation](https://github.com/bytesByHarsh/go-my-info/blob/master/docs/assets/swagger_1.png?raw=true)

![Swagger Documentation](https://github.com/bytesByHarsh/go-my-info/blob/master/docs/assets/swagger_2.png?raw=true)

![Swagger Documentation of Models](https://github.com/bytesByHarsh/go-my-info/blob/master/docs/assets/swagger_3.png?raw=true)

# Future Scope

- Encryption & Decryption Support
- Multiple APIs for different uses


## Setup

> Note: Make sure you have create a database in Postgres called `go_my_info`

```bash
go mod tidy
cd sql/schema
goose postgres postgres://<user_id>:<user_password>@localhost:5432/go_my_info up

cd ../..
go build ./cmd/main.go && ./main
```

## Dev Env

```bash
air
```
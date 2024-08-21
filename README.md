# Go My Info

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
goose postgres postgres://<user_id>:<user_password>@localhost:5432/rssAgg up
```
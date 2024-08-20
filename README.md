# Go My Info

### Setup

```bash
go install github.com/air-verse/air@latest
```

## Update Database

```bash
cd sql/schema
goose postgres postgres://<user_id>:<user_password>@localhost:5432/rssAgg up
```
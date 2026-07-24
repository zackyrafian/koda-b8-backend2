# Koda B8 Backend

REST API backend built with Go, Gin, PostgreSQL, and JWT authentication.

## Tech Stack

- **Go** + **Gin** — HTTP framework
- **PostgreSQL** — database (via pgx/v5)
- **JWT** — authentication
- **Swagger** — API documentation

## Project Structure

```
cmd/          entry point
internal/
  domain/     models & entities
  handler/    HTTP handlers
  middleware/ auth & other middleware
  repository/ database queries
  service/    business logic
  utils/      helpers
  libs/       shared libraries
```

## API Docs

Swagger UI tersedia di `/swagger/index.html` setelah server berjalan.

# go-web-demo

A minimal, production-style REST API in Go demonstrating clean layering (HTTP handlers → services → repositories), SQLite persistence, request validation, and unit testing with mocks.

## Highlights

- Gin for the HTTP API with clear routing and request models
- Layered architecture with separation of concerns (handlers, services, repositories)....
- SQLite3 database with migrations (pure SQL) with optional mock seed data
- Password hashing using bcrypt (and encoded with base64 for storage)
- Auth middleware using a fixed mock bearer token for protected routes
- Structured logging with zap
- Unit tests for service layer using testify and generated mocks
- Postman collection included for quick API testing

## Tech stack

- Web: github.com/gin-gonic/gin
- DB: SQLite (github.com/mattn/go-sqlite3)
- Logging: go.uber.org/zap
- Crypto: golang.org/x/crypto/bcrypt

## Project structure

```
cmd/web-demo/
    main.go                 # Entry point (bootstraps app and logging)
internal/app/
    app.go                  # Launch: DB connection and HTTP server
internal/http/
    server.go               # Routes and middleware wiring
    handler/                # HTTP handlers (users, posts)
    handler_model/          # Request DTOs with validation tags
    middleware/             # Auth (mock bearer token)
internal/services/        # Business logic (users, posts)
internal/repository/      # Persistence layer (users, posts)
internal/db/sqlite.go     # SQLite connection (+ PRAGMA foreign_keys)
migrations/               # Base schema (no mock data)
migrations-mock/          # Base schema + mock data
tests/                    # Service tests with mocks
demo.db                   # Local SQLite database file
```

## Getting started

Prerequisites:
- Go toolchain installed
- migrate CLI installed (https://github.com/golang-migrate/migrate)

1) Install dependencies (Go modules handle this automatically when you build/run).

2) Create the SQLite database via migrations (database path is ./demo.db as configured in code):

```bash
# Base schema
migrate -path migrations -database "sqlite3://demo.db" up

# Or: base schema + mock data
migrate -path migrations-mock -database "sqlite3://demo.db" up
```

3) Run the service (defaults to :8080):

```bash
go run ./cmd/web-demo
```

4) Health check:

```bash
curl --location 'http://localhost:8080/ping'
```

## Usage examples

Note on auth: protected routes require a mock bearer token. Obtain it via login or use the known value directly: `MOCK_VALID_JWT`.

### Users

- Create user

```bash
curl --location 'http://localhost:8080/user' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "username": "angelo",
        "email": "angelorodem@gmail.com",
        "password": "VeryNicePassw00rd!"
    }'
```

- Login (returns a mock token)

```bash
curl --location 'http://localhost:8080/user/login' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "email": "angelorodem@gmail.com",
        "password": "VeryNicePassw00rd!"
    }'
```

- Get user (requires bearer token)

```bash
curl --location --request GET 'http://localhost:8080/user' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer MOCK_VALID_JWT' \
    --data-raw '{
        "email": "angelorodem@gmail.com"
    }'
```

- Change username (requires bearer token)

```bash
curl --location --request PATCH 'http://localhost:8080/user' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer MOCK_VALID_JWT' \
    --data-raw '{
        "email": "angelorodem@gmail.com",
        "newUsername": "Angelus IV"
    }'
```

- Delete user (requires bearer token)

```bash
curl --location --request DELETE 'http://localhost:8080/user' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer MOCK_VALID_JWT' \
    --data-raw '{
        "email": "angelorodem@gmail.com"
    }'
```

### Posts

- Create post

```bash
curl --location 'http://localhost:8080/post' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "userEmail": "angelorodem@gmail.com",
        "title": "Welcome to the Blog!",
        "content": "Welcome to the Blog, I hope you have a nice time around!"
    }'
```

- Read post (public)

```bash
curl --location --request GET 'http://localhost:8080/post' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "id": 5
    }'
```

- Read all posts (public)

```bash
curl --location 'http://localhost:8080/post/all'
```

- Update post (requires bearer token and ownership)

```bash
curl --location --request PUT 'http://localhost:8080/post' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer MOCK_VALID_JWT' \
    --data-raw '{
        "id": 5,
        "userEmail": "angelorodem@gmail.com",
        "newTitle": "Welcomen!",
        "newContent": "Hope you are doing good!"
    }'
```

- Delete post (requires bearer token and ownership)

```bash
curl --location --request DELETE 'http://localhost:8080/post' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer MOCK_VALID_JWT' \
    --data-raw '{
        "id": 5,
        "userEmail": "angelorodem@gmail.com"
    }'
```

## Testing

Run the unit tests for the service layer:

```bash
go test ./tests/... -v
```

The tests use testify with mocks (generated with mockery) to validate service behavior independently of the database.

## Postman collection

A ready-to-use Postman collection is included: `Web-demo.postman_collection.json`.
Import it into Postman to try all endpoints quickly. Protected requests are preconfigured with the mock bearer token (`MOCK_VALID_JWT`). Update host/port if your service is not running on `localhost:8080`.

## Design notes

- Authentication is intentionally mocked for simplicity: the login endpoint verifies the bcrypt-hashed password and returns a fixed token (`MOCK_VALID_JWT`). Protected routes use a middleware that validates the bearer token against this constant. In a real service, you would issue/verify JWT or Paseto tokens and read user claims from them.
- Post ownership is checked in the service layer by resolving the `userEmail` to a user id and matching it against the post’s owner. With real tokens, you’d use the authenticated subject instead of passing `userEmail` in the request.
- SQLite is used for easy local setup. Foreign keys are enabled via PRAGMA.
- Request validation is done through Gin binding tags in the handler models.

## Roadmap / Ideas

- Replace mock auth with JWT/Paseto and proper claim checks
- Add environment-based configuration (flags/env, 12-factor app style)
- Containerize with Docker and add a Makefile

## License

This project is open source. See the `LICENSE` file for details.
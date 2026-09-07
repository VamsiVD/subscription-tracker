# Subscription Tracker

Small REST API for tracking subscriptions (like Netflix, gym, whatever you pay for). Written in Go, using `gin` for HTTP, `sqlx` + Postgres for db.

## Stack
- Go 1.27
- [gin](https://github.com/gin-gonic/gin) - HTTP router
- [sqlx](https://github.com/jmoiron/sqlx) + `lib/pq` - Postgres
- [godotenv](https://github.com/joho/godotenv) - load `.env`
- [google/uuid](https://github.com/google/uuid) - id generation

## Setup

1. Make Postgres db, make table `subscriptions` with column: `uuid`, `owner_id`, `name`, `type`, `amount`, `start_date`, `currency`.
2. Make `.env` file at repo root:
   ```
   DATABASE_URL=postgres://user:pass@localhost:5432/subscriptions?sslmode=disable
   ```
3. Run it:
   ```
   go run ./cmd/api
   ```
   Server start on `:8080` (gin default).

## API

| Method | Path | Body | What it do |
|---|---|---|---|
| POST | `/subscriptions` | `{ name, type, amount, start_date, currency, owner_id }` | Create subscription |
| GET | `/subscriptions/:owner_id` | - | List all subscription for owner |
| PATCH | `/subscriptions/:uuid` | `{ name, type, amount, start_date, currency }` | Update subscription |
| DELETE | `/subscriptions/:uuid` | - | Delete subscription |

`type` is one of `weekly`, `monthly`, `yearly`. `amount` is integer (cent, probably).

## Project layout
```
cmd/api/main.go              entry point, wire everything together
internal/config               load env config
internal/db                   db connect
internal/server                gin router
internal/subscriptions         model, repository, handler for subscription CRUD
```

## Status
Core CRUD live. No auth yet, no test yet - next step for this project.

---


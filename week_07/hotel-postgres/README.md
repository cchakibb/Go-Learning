# Hotel Rooms API

A RESTful hotel rooms API built with Go and PostgreSQL, demonstrating 
persistent CRUD operations over HTTP. Data is stored in a PostgreSQL 
database running in Docker, accessed via raw SQL (no ORM).

## Architecture

HTTP Client (curl / browser)

│  HTTP requests on :8080

▼

Go HTTP Server (local)

├── handlers.go  — routes requests, encodes/decodes JSON

├── store.go     — executes raw SQL via database/sql + pgx driver

│  TCP on localhost:5433

▼

PostgreSQL 16 (Docker container)

└── volume: data persisted on host disk

The Go server runs locally. PostgreSQL runs in a Docker container with 
its data directory mounted as a named volume — data survives container 
restarts and removals.


## Prerequisites

- [Go 1.26+](https://go.dev/dl/)
- [Docker](https://docs.docker.com/get-docker/) and Docker Compose
- `psql` (PostgreSQL client) — `sudo apt install postgresql-client`
- `jq` (optional, for readable JSON output) — `sudo apt install jq`

## Getting Started

### 1. Start the database

```bash
docker compose up -d
docker compose ps   # verify the container is running
```


### 2. Create the table

Connect to the database:

```bash
psql -h localhost -p 5433 -U hotel -d hoteldb
# password: secret
```

Create the rooms table:

```sql
CREATE TABLE rooms (
    id    TEXT PRIMARY KEY,
    name  TEXT NOT NULL,
    price FLOAT NOT NULL
);
```

You can also insert sample data manually:

```sql
INSERT INTO rooms (id, name, price) VALUES ('101', 'Suite', 150);
```

Or skip this step and use the API endpoints below to populate the database.

### 3. Start the server

```bash
go run .
```

The server starts on `http://localhost:8080`.

## Database Schema

```sql
CREATE TABLE rooms (
    id    TEXT PRIMARY KEY,   -- unique identifier, cannot be null
    name  TEXT NOT NULL,      -- room name
    price FLOAT NOT NULL      -- price per night
);
```

`PRIMARY KEY` on `id` means: unique, non-null, and automatically indexed.
PostgreSQL uses a B-tree index on `id` for fast lookups — no full table
scan needed when searching by ID.

| Column | SQL Type | Go Type  | Constraint        |
|--------|----------|----------|-------------------|
| id     | TEXT     | string   | PRIMARY KEY       |
| name   | TEXT     | string   | NOT NULL          |
| price  | FLOAT    | float64  | NOT NULL          |

## API Endpoints

### GET /rooms
Returns all rooms as a JSON array.

```bash
curl -s http://localhost:8080/rooms | jq .
```

### POST /rooms
Creates a new room. Sends a JSON body with `id`, `name`, and `price`.
Returns `201 Created` on success, `409 Conflict` if the ID already exists.

```bash
curl -s -X POST http://localhost:8080/rooms \
  -H "Content-Type: application/json" \
  -d '{"id":"101","name":"Suite","price":150}' | jq .
```

### GET /rooms/{id}
Returns a single room by ID.
Returns `404 Not Found` if the room does not exist.

```bash
curl -s http://localhost:8080/rooms/101 | jq .
```

### DELETE /rooms/{id}
Deletes a room by ID.
Returns `204 No Content` on success, `404 Not Found` if not found.

```bash
curl -s -o /dev/null -w "%{http_code}" \
  -X DELETE http://localhost:8080/rooms/101
```

### PATCH /rooms/{id}/price
Updates the price of a room. Uses a database transaction to ensure
atomicity: the room is verified to exist before the price is updated.
Returns the updated room as JSON, or `404` if the room does not exist.

```bash
curl -s -X PATCH http://localhost:8080/rooms/101/price \
  -H "Content-Type: application/json" \
  -d '{"price":299.99}' | jq .
```

## Environment

The connection string used by the Go server:
host=localhost port=5433 user=hotel password=secret dbname=hoteldb sslmode=disable

| Variable | Value     | Note                                      |
|----------|-----------|-------------------------------------------|
| host     | localhost | PostgreSQL runs in Docker on your machine |
| port     | 5433      | host port mapped to container port 5432   |
| user     | hotel     | defined in docker-compose.yml             |
| password | secret    | defined in docker-compose.yml             |
| dbname   | hoteldb   | defined in docker-compose.yml             |
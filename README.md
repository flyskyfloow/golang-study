# golang-study — Todo API

A simple REST API for managing Todo items, built with [Go](https://go.dev/) and the [Gin](https://github.com/gin-gonic/gin) web framework.

## Requirements

- Go 1.25+
- Docker (optional, for containerised deployment)

## Running locally

```bash
# Download dependencies
go mod download

# Start the server (listens on :8080)
go run .
```

## API endpoints

| Method | Path              | Description          |
|--------|-------------------|----------------------|
| POST   | /api/v1/todos     | Create a new todo    |
| GET    | /api/v1/todos     | List all todos       |
| PUT    | /api/v1/todos/:id | Update a todo by ID  |
| DELETE | /api/v1/todos/:id | Delete a todo by ID  |

### Example requests

**Create**
```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Buy groceries"}'
```

**List**
```bash
curl http://localhost:8080/api/v1/todos
```

**Update**
```bash
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Buy groceries","completed":true}'
```

**Delete**
```bash
curl -X DELETE http://localhost:8080/api/v1/todos/1
```

## Running with Docker

```bash
# Build the image
docker build -t todo-api .

# Run the container
docker run -p 8080:8080 todo-api
```

## Running tests

```bash
go test ./...
```

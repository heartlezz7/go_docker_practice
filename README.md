# Go Docker Practice

A simple REST API built with Go and the Gin framework, containerized with Docker using a multi-stage build.

## Tech Stack

- **Go** 1.22
- **Gin** — HTTP web framework
- **Docker** — multi-stage build with Alpine base image
- **Docker Compose** — local orchestration with MongoDB

## API Endpoints

| Method | Path     | Description          |
|--------|----------|----------------------|
| GET    | `/ping`  | Health check         |
| POST   | `/login` | Echo user JSON body  |

### POST /login — Request Body

```json
{
  "Name": "John",
  "Email": "john@example.com"
}
```

## Getting Started

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/) installed

### Run with Docker

```bash
# Build the image
docker build -t my-app .

# Start the container
docker run -d -p 8080:8080 --name my-app my-app

# Check it's running
docker ps

# Stop the container
docker stop my-app
```

### Run with Docker Compose

```bash
# Start app + MongoDB
docker compose up -d

# Stop everything
docker compose down
```

### Run Locally (without Docker)

```bash
# Copy and configure environment
touch .env   # set PORT=8080

go run main.go
```

## Project Structure

```
.
├── main.go           # Application entry point & routes
├── dockerfile        # Multi-stage Docker build
├── docker-compose.yml
├── go.mod / go.sum
└── README.md
```

## Docker Image Details

The Dockerfile uses a two-stage build:

1. **Builder stage** — compiles the Go binary (`golang:1.22`)
2. **Runtime stage** — runs the binary on a minimal `alpine:latest` image

This keeps the final image small by excluding the Go toolchain.
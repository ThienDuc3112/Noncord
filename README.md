# Noncord

A modern, Discord-inspired real-time chat application built with a microservices architecture, featuring WebSocket support for real-time messaging and event-driven communication.

> **Architecture inspired by [sklinkert/go-ddd](https://github.com/sklinkert/go-ddd).**

## Overview

Noncord is a full-stack chat application that implements server-based communication with channels, direct messaging, user management, and real-time updates. The project follows clean architecture principles and uses an event-driven approach with the transactional outbox pattern.

## Features

### Implemented

- User registration and authentication  
- Server creation and management  
- Channel creation and management  
- Real-time messaging via WebSocket  
- Server invitations  
- Membership management  
- Event-driven architecture with outbox pattern  

### In Progress

- Role-based permissions  
- Direct messaging groups  
- File attachments  
- User notifications  
- Advanced permission system  

## Tech Stack

### Backend

- **Language**: Go 1.24.3  
- **Web Framework**: Chi (v5.2.1)  
- **Database**: PostgreSQL 17 (pgx v5.7.5)  
- **Database Tooling**: sqlc  
- **Message Queue**: RabbitMQ 4.2.0  
- **WebSocket**: Gorilla WebSocket (v1.5.3)  
- **Authentication**: JWT (golang-jwt/jwt v5.2.2)  
- **API Documentation**: Swagger/OpenAPI with swaggo  

### Frontend

- **Framework**: Next.js 15.5.4 with React 19  
- **Language**: TypeScript 5  
- **Styling**: Tailwind CSS 4  
- **UI Components**: Radix UI primitives  
- **State Management**: TanStack Query (v5.90.2)  
- **Forms**: React Hook Form + Zod  
- **HTTP Client**: Axios  

## Project Structure

```bash
Noncord/
├── backend/
│   ├── cmd/              # Application entry points
│   │   ├── api/          # REST API service
│   │   ├── ws/           # WebSocket service
│   │   └── relayer/      # Event relayer service
│   ├── internal/         # Backend architecture (see backend/README.md)
│   ├── docs/             # API documentation
│   └── go.mod
├── frontend/
│   ├── app/              # Next.js app directory
│   ├── components/       # React components
│   ├── lib/              # Utility libraries
│   └── package.json
└── docker-compose.yml    # Docker orchestration
````

For a detailed description of the backend architecture, see:
**[`backend/README.md`](backend/README.md)**

---

## Getting Started

### Using Docker Compose (Recommended)

1. Clone the repository:

```bash
git clone <repository-url>
cd Noncord
```

2. Start all services:

```bash
docker-compose up
```

This will start:

* PostgreSQL database on port `6543`
* RabbitMQ on ports `5672` (AMQP) and `15672` (Management UI)
* WebSocket service on port `9999`
* REST API service on port `8888`
* Frontend application on port `6969`

3. Access the application:

* **Frontend**: [http://localhost:6969](http://localhost:6969)
* **API**: [http://localhost:8888/api/v1](http://localhost:8888/api/v1)
* **API Documentation**: [http://localhost:8888/api/v1/docs/](http://localhost:8888/api/v1/docs/)
* **RabbitMQ Management**: [http://localhost:15672](http://localhost:15672) (username: `noncord`, password: `noncord`)

---

## Local Development

### Backend (API, WebSocket, Relayer)

Backend-specific details live in [`backend/README.md`](backend/README.md), but the short version:

1. Set environment variables:

```bash
export DB_URI="postgres://noncord:password@localhost:6543/noncord?sslmode=disable"
export AMQP_URI="amqp://noncord:noncord@localhost:5672/"
export PORT="8888"
export SECRET="your-jwt-secret"
```

2. Start database and message queue:

```bash
docker-compose up db mq
```

3. Run migrations and generate code:

```bash
cd backend
go install github.com/pressly/goose/v3/cmd/goose@latest
goose -dir internal/infra/db/sql/migration postgres "$DB_URI" up
sqlc generate
```

4. Run services:

```bash
# From backend/
go run cmd/api/main.go      # REST API
go run cmd/ws/main.go       # WebSocket service
go run cmd/relayer/main.go  # Outbox relayer
```

### Frontend

```bash
cd frontend
npm install
```

Create `.env`:

```env
API_URL=http://localhost:8888/api/v1
NEXT_PUBLIC_BACKEND_URL=http://localhost:8888/api/v1
```

Run dev server:

```bash
npm run dev
```

---

## API Documentation

When the API service is running, Swagger docs are available at:

* [http://localhost:8888/api/v1/docs/](http://localhost:8888/api/v1/docs/)

---

## Contributing

This is an active development project.
Check `backend/TODO.md` (or `TODO.md` in the backend directory) for current development tasks and priorities.

---

## License

To be decided.

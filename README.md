# Noncord

A modern, Discord-inspired real-time chat application built with a microservices architecture, featuring WebSocket support for real-time messaging and event-driven communication.

> **Architecture inspired by [sklinkert/go-ddd](https://github.com/sklinkert/go-ddd).**

## Overview

Noncord is a full-stack chat application that implements server-based communication with channels, direct messaging, user management, and real-time updates. The project follows clean architecture principles and uses an event-driven approach with the transactional outbox pattern.

## Architecture

The application consists of multiple services working together:

### Backend Services

1. **API Service** - RESTful API handling HTTP requests for user, server, channel, and message management  
2. **WebSocket Service** - Real-time bidirectional communication for live updates  
3. **Relayer Service** - Event relayer that processes outbox events and publishes them to RabbitMQ  

### Infrastructure Components

- **PostgreSQL Database** - Primary data store for all application data  
- **RabbitMQ** - Message broker for event-driven communication between services  

### Frontend

- **Next.js Application** - Modern React-based UI with server-side rendering  

## Technology Stack

### Backend
- **Language**: Go 1.24.3  
- **Web Framework**: Chi (v5.2.1)  
- **Database**: PostgreSQL 17 with pgx driver (v5.7.5)  
- **Database Tooling**: sqlc for type-safe SQL queries  
- **Message Queue**: RabbitMQ 4.2.0  
- **WebSocket**: Gorilla WebSocket (v1.5.3)  
- **Authentication**: JWT tokens with golang-jwt/jwt (v5.2.2)  
- **API Documentation**: Swagger/OpenAPI with swaggo  
- **Development**: Air for hot reloading  

### Frontend
- **Framework**: Next.js 15.5.4 with React 19  
- **Language**: TypeScript 5  
- **Styling**: Tailwind CSS 4  
- **UI Components**: Radix UI primitives  
- **State Management**: TanStack Query (v5.90.2)  
- **Forms**: React Hook Form with Zod validation  
- **HTTP Client**: Axios  

## Project Structure

```bash
Noncord/
├── backend/
│   ├── cmd/              # Application entry points
│   │   ├── api/          # REST API service
│   │   ├── ws/           # WebSocket service
│   │   └── relayer/      # Event relayer service
│   ├── internal/
│   │   ├── application/  # Application layer (services, commands, queries)
│   │   ├── domain/       # Domain layer (entities, events, repositories)
│   │   ├── infra/        # Infrastructure layer (database, message queue)
│   │   ├── interface/    # Interface layer (REST, WebSocket)
│   │   └── processes/    # Background processes
│   ├── docs/             # API documentation
│   └── go.mod
├── frontend/
│   ├── app/              # Next.js app directory
│   ├── components/       # React components
│   ├── lib/              # Utility libraries
│   └── package.json
└── docker-compose.yml    # Docker orchestration
````

## Domain Model

The application manages the following core entities:

* **Users** - User accounts with authentication and profiles
* **Servers** - Community spaces that contain channels
* **Channels** - Communication channels within servers
* **Messages** - Text messages in channels or direct messages
* **Memberships** - User membership in servers with roles
* **Invitations** - Server invitation system
* **Roles** - Permission-based access control
* **Sessions** - User authentication sessions

## Prerequisites

* Docker and Docker Compose
* Go 1.24.3 or higher (for local development)
* Node.js and npm (for local development)

## Getting Started

### Using Docker Compose (Recommended)

1. Clone the repository:

```bash
git clone https://github.com/ThienDuc3112/Noncord
cd Noncord
```

2. Start all services:

```bash
docker-compose up
```

This will start:

* PostgreSQL database on port 6543
* RabbitMQ on ports 5672 (AMQP) and 15672 (Management UI)
* WebSocket service on port 9999
* REST API service on port 8888
* Frontend application on port 6969

3. Access the application:

* Frontend: [http://localhost:6969](http://localhost:6969)
* API: [http://localhost:8888/api/v1](http://localhost:8888/api/v1)
* API Documentation: [http://localhost:8888/api/v1/docs/](http://localhost:8888/api/v1/docs/)
* RabbitMQ Management: [http://localhost:15672](http://localhost:15672) (username: noncord, password: noncord)

### Local Development

#### Backend

1. Set up environment variables:

```bash
export DB_URI="postgres://noncord:password@localhost:6543/noncord?sslmode=disable"
export AMQP_URI="amqp://noncord:noncord@localhost:5672/"
export PORT="8888"
export SECRET="your-jwt-secret"
```

2. Start the database and message queue:

```bash
docker-compose up db mq
```

3. Run database migrations and generate code:

```bash
cd backend
# Install goose for migrations
go install github.com/pressly/goose/v3/cmd/goose@latest
# Run migrations
goose -dir internal/infra/db/sql/migration postgres "$DB_URI" up
# Generate sqlc code
sqlc generate
```

4. Run services:

```bash
# API Service
go run cmd/api/main.go

# WebSocket Service
go run cmd/ws/main.go

# Relayer Service
go run cmd/relayer/main.go
```

#### Frontend

1. Install dependencies:

```bash
cd frontend
npm install
```

2. Create `.env` file with:

```env
API_URL=http://localhost:8888/api/v1
NEXT_PUBLIC_BACKEND_URL=http://localhost:8888/api/v1
```

3. Run development server:

```bash
npm run dev
```

## Features

### Implemented

* User registration and authentication
* Server creation and management
* Channel creation and management
* Real-time messaging via WebSocket
* Server invitations
* Membership management
* Event-driven architecture with outbox pattern

### In Progress

* Role-based permissions
* Direct messaging groups
* File attachments
* User notifications
* Advanced permission system

## API Documentation

The REST API is documented using Swagger/OpenAPI. When running the API service, access the documentation at:

[http://localhost:8888/api/v1/docs/](http://localhost:8888/api/v1/docs/)

## Development Tools

* **Air**: Hot reloading for Go applications
* **sqlc**: Type-safe SQL code generation
* **goose**: Database migration tool
* **Swagger**: API documentation generation

## Environment Variables

### Backend Services

* `DB_URI`: PostgreSQL connection string
* `AMQP_URI`: RabbitMQ connection string
* `PORT`: Service port number
* `SECRET`: JWT signing secret
* `FRONTEND_URL`: Frontend URL for CORS
* `ENVIRONMENT`: Development/production mode

### Frontend

* `API_URL`: Internal API URL (for SSR)
* `NEXT_PUBLIC_BACKEND_URL`: Public API URL (for client-side)

## Contributing

This is an active development project. Check the TODO.md file in the backend directory for current development tasks and priorities.

## Architecture Patterns

* **Clean Architecture**: Separation of concerns with domain, application, infrastructure, and interface layers
* **Event Sourcing**: Domain events captured and published
* **Transactional Outbox Pattern**: Reliable event publishing with database transactions
* **CQRS**: Command and Query separation in the application layer
* **Microservices**: Independent, scalable services

## License

To be decided

# Noncord Backend

The Noncord backend powers the API, WebSocket, and event relayer services for the application. It follows clean architecture and Domain-Driven Design, with an event-driven, outbox-based integration model.

> **Architecture inspired by [sklinkert/go-ddd](https://github.com/sklinkert/go-ddd).**

For an overview of the full project (frontend + backend, how to run everything), see the root [`README.md`](../README.md).

---

## Services & Infrastructure

### Backend Services

1. **API Service (`cmd/api`)**  
   RESTful HTTP API for:
   - User registration and authentication  
   - Server and channel management  
   - Membership and invitation management  
   - Message CRUD (where applicable)

2. **WebSocket Service (`cmd/ws`)**  
   - Real-time bidirectional communication using Gorilla WebSocket  
   - Subscriptions to server/channel events  
   - Pushes new messages and updates to connected clients

3. **Relayer Service (`cmd/relayer`)**  
   - Reads **outbox** events from PostgreSQL  
   - Publishes them to RabbitMQ  
   - Acts as a bridge between domain events and the messaging infrastructure

### Infrastructure Components

- **PostgreSQL 17**  
  Primary data store for all application data, including outbox tables.

- **RabbitMQ 4.2.0**  
  Message broker used for event-driven communication between services.

---

## Layered Architecture

The backend is structured using clean architecture:

```bash
backend/
├── cmd/
│   ├── api/          # API service entrypoint (main.go, wiring)
│   ├── ws/           # WebSocket service entrypoint
│   └── relayer/      # Outbox relayer entrypoint
├── internal/
│   ├── application/  # Application layer (services, commands, queries)
│   ├── domain/       # Domain layer (entities, value objects, events, repos)
│   ├── infra/        # Infrastructure (Postgres, RabbitMQ, logging, etc.)
│   ├── interface/    # HTTP/WS handlers, DTOs, routing
│   └── processes/    # Background processes, workers
└── docs/             # Generated API docs (Swagger)
````

### Domain Layer (`internal/domain`)

* **Entities**

  * `User`
  * `Server`
  * `Channel`
  * `Message`
  * `Membership`
  * `Invitation`
  * `Role`
  * `Session`

* **Domain Events**
  Events such as:

  * `server.created`
  * `channel.created`
  * `message.created`
  * `membership.joined`
  * `invitation.created`
    Each event has an associated schema version and payload.

* **Repositories (Interfaces)**
  Abstractions for persistence (e.g., `UserRepository`, `ServerRepository`, `MessageRepository`), implemented in the infrastructure layer.

### Application Layer (`internal/application`)

* **Commands**

  * Mutating use cases, like `CreateServer`, `JoinServer`, `SendMessage`, etc.
  * Use domain entities + repositories + domain services.
  * Raise domain events on state changes.

* **Queries**

  * Read-only use cases, like `GetServer`, `ListChannels`, `GetMessages`.
  * May use optimized read models if needed (CQRS).

* **Services**

  * Orchestrate domain logic and enforce application-level invariants.
  * Use repositories and other services (e.g., permission checks, invitations).

### Interface Layer (`internal/interface`)

* **HTTP Handlers** (API)

  * Map HTTP routes and DTOs to application commands/queries.
  * Validation using appropriate libraries (e.g., JSON parsing, struct tags).

* **WebSocket Handlers** (WS service)

  * Connection management and routing.
  * Subscribes users to servers/channels.
  * Writes outgoing messages/events to sockets.

### Infrastructure Layer (`internal/infra`)

* **Database (PostgreSQL + pgx)**

  * `sqlc`-generated repositories for type-safe SQL.
  * Migration scripts under `internal/infra/db/sql/migration`.

* **Message Broker (RabbitMQ)**

  * Publishes domain events (via outbox/relayer).
  * Subscribes to events for cross-service communication.

* **Logging, Config, Middleware**

  * Common concerns implemented here.

---

## Architecture Patterns

* **Clean Architecture**

  * Domain and application layers are independent of frameworks and infrastructure.
  * Infrastructure and interfaces depend on domain/application, never the other way around.

* **Domain-Driven Design (DDD)**

  * Clear aggregates: servers, channels, memberships, etc.
  * Domain events model important state transitions.

* **CQRS**

  * Commands (writes) and queries (reads) separated at the application layer.
  * Queries can use tailored read models without affecting write models.

* **Event-Driven Architecture**

  * Domain events raised in the domain layer.
  * Persisted in an **outbox** table as part of the same DB transaction.

* **Transactional Outbox Pattern**

  * On a successful transaction, new events are inserted into an outbox table.
  * Relayer service periodically claims outbox rows, publishes them to RabbitMQ, and marks them as processed.
  * Ensures at-least-once delivery and avoids dual-write inconsistencies.

---

## Domain Model (Conceptual)

The backend manages these core entities:

* **Users**
  Authenticated accounts with profiles.

* **Servers**
  Community spaces that group users and channels.

* **Channels**
  Conversation spaces within servers.

* **Messages**
  Text messages (and in future, attachments) in channels or direct messages.

* **Memberships**
  Relationship between users and servers, including assigned roles.

* **Invitations**
  Token-based server invitation system (with expiry, join limits, etc.).

* **Roles**
  Permission-based access control units (e.g., admin, moderator, member).

* **Sessions**
  Authentication sessions associated with JWTs or tokens.

---

## Backend Tech Stack

* **Language**: Go 1.24.3
* **HTTP Framework**: Chi (v5.2.1)
* **Database**: PostgreSQL 17 + pgx (v5.7.5)
* **Database Tooling**: sqlc, goose
* **Message Queue**: RabbitMQ 4.2.0
* **WebSocket**: Gorilla WebSocket (v1.5.3)
* **Auth**: JWT via golang-jwt/jwt (v5.2.2)
* **API Docs**: swaggo/Swagger
* **Dev Tooling**: Air for hot reloading

---

## Environment Variables (Backend)

* `DB_URI`
  PostgreSQL connection string, e.g.
  `postgres://noncord:password@localhost:6543/noncord?sslmode=disable`

* `AMQP_URI`
  RabbitMQ connection string, e.g.
  `amqp://noncord:noncord@localhost:5672/`

* `PORT`
  Port for the running service (API, WS, etc.).

* `SECRET`
  JWT signing secret.

* `FRONTEND_URL`
  Used for CORS configuration.

* `ENVIRONMENT`
  Environment name, e.g. `development`, `production`.

---

## Running the Backend (Dev)

Short version (for details, see root `README.md`):

```bash
# From project root
docker-compose up db mq
```

Set env vars:

```bash
export DB_URI="postgres://noncord:password@localhost:6543/noncord?sslmode=disable"
export AMQP_URI="amqp://noncord:noncord@localhost:5672/"
export PORT="8888"
export SECRET="your-jwt-secret"
```

Run migrations and generate sqlc code:

```bash
cd backend
go install github.com/pressly/goose/v3/cmd/goose@latest
goose -dir internal/infra/db/sql/migration postgres "$DB_URI" up
sqlc generate
```

Start services:

```bash
# API
go run cmd/api/main.go

# WebSocket
go run cmd/ws/main.go

# Relayer
go run cmd/relayer/main.go
```

---

## API Documentation

Once the API is running, Swagger docs are exposed at:

* [http://localhost:8888/api/v1/docs/](http://localhost:8888/api/v1/docs/)

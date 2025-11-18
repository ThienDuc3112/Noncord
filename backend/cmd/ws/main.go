package main

import (
	"backend/internal/application/services"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres"
	rabbitmq "backend/internal/infra/rabbitMQ"
	"backend/internal/interface/ws"
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, os.Interrupt)
	defer stop()

	conn, err := pgxpool.New(ctx, os.Getenv("DB_URI"))

	if err != nil {
		cancel()
		log.Fatalf("Cannot connect to db: %v", err)
	}

	uow := postgres.NewBaseUoW(conn)

	permissionService := services.NewPermissionService(postgres.NewScopedUoW(uow, func(rb repositories.RepoBundle) services.PermissionRepos { return rb }))

	rabbitMQConn, err := amqp091.Dial(os.Getenv("AMQP_URI"))
	if err != nil {
		cancel()
		log.Fatalf("Cannot connect to rabbitMQ: %v", err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	eventSub, err := rabbitmq.NewRMQEventConsumer(ctx, rabbitMQConn, logger, "websocket", "noncord.event", true)
	if err != nil {
		cancel()
		log.Fatalf("Cannot connect to rabbitMQ: %v", err)
	}
	defer eventSub.Close(ctx)

	wsHub, err := ws.NewHub(ctx, permissionService, eventSub, logger)
	if err != nil {
		cancel()
		log.Fatalf("Cannot connect to rabbitMQ: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Connect("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(wsHub, w, r)
	})

	log.Printf("listening on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), r))
}

package main

import (
	"backend/internal/application/services"
	"backend/internal/domain/repositories"
	"backend/internal/infra/cache/inmemcache"
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
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, os.Interrupt)
	defer stop()

	pgPool, err := pgxpool.New(ctx, os.Getenv("DB_URI"))

	if err != nil {
		cancel()
		log.Fatalf("Cannot connect to db: %v", err)
	}

	uow := postgres.NewBaseUoW(pgPool)

	visiblityQueries := services.NewVisibilityQueries(postgres.NewScopedUoW(uow, func(rb repositories.RepoBundle) services.PermissionRepos { return rb }))
	authService := services.NewAuthService(postgres.NewScopedUoW(uow, func(rb repositories.RepoBundle) services.AuthRepos { return rb }), os.Getenv("SECRET"))

	rabbitMQConn, err := amqp091.Dial(os.Getenv("AMQP_URI"))
	if err != nil {
		cancel()
		log.Fatalf("Cannot connect to rabbitMQ: %v", err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	slog.SetDefault(logger)

	userResolver := postgres.NewPGNicknameResolver(pgPool)
	cacheStore := inmemcache.NewInMemoryCache(15*time.Minute, 2*time.Minute)

	eventSub, err := rabbitmq.NewRMQEventSubscriber(ctx, rabbitMQConn, "websocket", "noncord.event", true)
	if err != nil {
		cancel()
		log.Fatalf("Cannot connect to rabbitMQ: %v", err)
	}
	defer eventSub.Close()

	wsHub, err := ws.NewHub(ctx, visiblityQueries, eventSub, cacheStore, userResolver)
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

	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(wsHub, authService, w, r)
	})

	log.Printf("listening on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), r))
}

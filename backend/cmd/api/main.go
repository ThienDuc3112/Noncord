package main

import (
	_ "backend/docs"
	"backend/internal/application/services"
	"backend/internal/application/workers"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres"
	rabbitmq "backend/internal/infra/rabbitMQ"
	"backend/internal/interface/rest"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rabbitmq/amqp091-go"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title			Noncord API
// @version		1.0
// @description	This is the api for Noncord
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, os.Interrupt)
	defer stop()

	pgPool, err := pgxpool.New(ctx, os.Getenv("DB_URI"))
	if err != nil {
		log.Fatalf("Cannot connect to db: %v", err)
	}

	rabbitMQConn, err := amqp091.Dial(os.Getenv("AMQP_URI"))
	if err != nil {
		log.Fatalf("Cannot connect to rabbitMQ: %v", err)
	}

	eventSub, err := rabbitmq.NewRMQEventSubscriber(ctx, rabbitMQConn, "api_workers", "noncord.event", true)
	if err != nil {
		log.Fatalf("Cannot create a new event sub: %v", err)
	}

	uow := postgres.NewBaseUoW(pgPool)

	// ---------- Services ----------
	authService := services.NewAuthService(postgres.NewScopedUoW(uow, func(rb repositories.RepoBundle) services.AuthRepos {
		return rb
	}), os.Getenv("SECRET"))
	serverService := services.NewServerService(postgres.NewScopedUoW(uow, func(rb repositories.RepoBundle) services.ServerRepos { return rb }))
	invitationService := services.NewInvitationService(postgres.NewScopedUoW(uow, func(rb repositories.RepoBundle) services.InvitationRepos { return rb }))
	membershipService := services.NewMemberService(postgres.NewScopedUoW(uow, func(rb repositories.RepoBundle) services.MemberRepos { return rb }))
	channelService := services.NewChannelService(postgres.NewScopedUoW(uow, func(rb repositories.RepoBundle) services.ChannelRepos { return rb }))
	messageService := services.NewMessageService(postgres.NewScopedUoW(uow, func(rb repositories.RepoBundle) services.MessageRepos { return rb }))

	// ---------- Queries ----------
	serverQueries := postgres.NewPGServerQueries(pgPool)
	inviteQueries := services.NewInvitationQueries(postgres.NewScopedUoW(uow, func(rb repositories.RepoBundle) services.InvitationRepos { return rb }))
	messageQueries := postgres.NewPGMessageQueries(pgPool)
	channelQueries := services.NewChannelQueries(postgres.NewScopedUoW(uow, func(rb repositories.RepoBundle) services.ChannelRepos { return rb }))
	userQueries := postgres.NewPGUserQueries(pgPool)

	if err = workers.NewWorker(messageService, eventSub); err != nil {
		log.Fatalf("Cannot attach workers to event sub: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Route("/api/v1", func(r chi.Router) {
		docsHandler := httpSwagger.Handler(httpSwagger.URL(fmt.Sprintf("http://localhost:%v/api/v1/docs/doc.json", port)))
		r.Get("/docs/*", docsHandler)

		rest.NewAuthController(authService).RegisterRoute(r)
		rest.NewServerController(authService, serverService, serverQueries, invitationService, inviteQueries).RegisterRoute(r)
		rest.NewInvitationController(serverQueries, authService, invitationService, inviteQueries, membershipService).RegisterRoute(r)
		rest.NewMessageController(messageService, messageQueries, authService).RegisterRoute(r)
		rest.NewChannelController(authService, channelService, channelQueries).RegisterRoute(r)
		rest.NewUserController(authService, userQueries).RegisterRoute(r)
	})

	log.Printf("listening on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), r))
}

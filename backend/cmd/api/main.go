package main

import (
	_ "backend/docs"
	"backend/internal/application/services"
	"backend/internal/infra/db/postgres"
	"backend/internal/interface/api/rest"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/swaggo/http-swagger/v2"
)

// @title			Noncord API
// @version		1.0
// @description	This is the api for Noncord
func main() {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DB_URI"))

	if err != nil {
		log.Fatalf("Cannot connect to db: %v", err)
	}

	userRepo := postgres.NewPGUserRepo(conn)
	sessionRepo := postgres.NewPGSessionRepo(conn)
	serverRepo := postgres.NewPGServerRepo(conn)
	invitationRepo := postgres.NewPGInvitationRepo(conn)
	memberRepo := postgres.NewPGMemberRepo(conn)

	authService := services.NewAuthService(userRepo, sessionRepo, conn, os.Getenv("SECRET"))
	serverService := services.NewServerService(serverRepo, memberRepo)
	invitationService := services.NewInvitationService(serverRepo, invitationRepo)
	membershipService := services.NewMemberService(memberRepo, invitationRepo, serverRepo, userRepo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
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

	r.Route("/api/v1", func(r chi.Router) {
		docsHandler := httpSwagger.Handler(httpSwagger.URL(fmt.Sprintf("http://localhost:%v/api/v1/docs/doc.json", port)))
		r.Get("/docs/*", docsHandler)

		rest.NewAuthController(authService).RegisterRoute(r)
		rest.NewServerController(serverService, authService, invitationService).RegisterRoute(r)
		rest.NewInvitationController(serverService, authService, invitationService, membershipService).RegisterRoute(r)
	})

	log.Printf("listening on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), r))
}

package main

import (
	_ "backend/docs"
	"backend/internal/application/services"
	"backend/internal/infra/db/postgres"
	"backend/internal/interface/api/rest"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/swaggo/http-swagger/v2"
)

func main() {
	conn, err := sql.Open("pgx", os.Getenv("DB_URI"))
	if err != nil {
		log.Fatalf("Cannot connect to db: %v", err)
	}

	userRepo := postgres.NewPGUserRepo(conn)
	sessionRepo := postgres.NewPGSessionRepo(conn)

	authService := services.NewAuthService(userRepo, sessionRepo, conn, os.Getenv("SECRET"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(fmt.Sprintf("http://localhost:%v/api/v1/docs/doc.json", port))))

		rest.NewAuthController(authService).RegisterRoute(r)
	})

	log.Printf("listening on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), r))
}

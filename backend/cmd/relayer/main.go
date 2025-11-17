package main

import (
	"backend/internal/infra/db/postgres"
	rabbitmq "backend/internal/infra/rabbitMQ"
	"backend/internal/processes/relayer"
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	pgxConn, err := pgxpool.New(ctx, os.Getenv("DB_URI"))
	if err != nil {
		log.Fatalf("Cannot connect to db: %v", err)
	}

	amqpConn, err := amqp091.Dial(os.Getenv("AMQP_URI"))
	if err != nil {
		log.Fatalf("Cannot connect to rabbitMQ: %v", err)
	}
	mq := rabbitmq.NewRMQEventPublisher(amqpConn, slog.Default())
	outboxReader := postgres.NewPGOutboxReader(pgxConn)
	relayer := relayer.New(slog.Default(), outboxReader, mq, relayer.Config{
		BatchSize:    100,
		StaleAfter:   time.Minute,
		PollInterval: 100 * time.Millisecond,
		Topic:        "noncord events",
	})

	if err = relayer.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

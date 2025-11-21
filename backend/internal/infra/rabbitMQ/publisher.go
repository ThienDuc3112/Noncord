package rabbitmq

import (
	"backend/internal/application/ports"
	"context"
	"log/slog"

	amqp "github.com/rabbitmq/amqp091-go"
)

const EXCHANGE_NAME = "noncord.event"

type RMQEventPublisher struct {
	conn *amqp.Connection
}

func NewRMQEventPublisher(conn *amqp.Connection) ports.EventPublisher {
	return &RMQEventPublisher{conn}
}

func (mq *RMQEventPublisher) Publish(ctx context.Context, msg ports.EventMessage) error {
	c, err := mq.conn.Channel()
	if err != nil {
		return err
	}

	if err = c.ExchangeDeclare(EXCHANGE_NAME, "topic", true, false, false, false, nil); err != nil {
		return err
	}

	if err = c.Publish(EXCHANGE_NAME, msg.EventType, false, false, amqp.Publishing{
		Headers:     msg.Headers,
		ContentType: "application/json",
		Body:        msg.Payload,
	}); err != nil {
		return err
	}

	return nil
}

func (mq *RMQEventPublisher) Close(ctx context.Context) error {
	slog.Info("Closing publisher")
	return mq.conn.Close()
}

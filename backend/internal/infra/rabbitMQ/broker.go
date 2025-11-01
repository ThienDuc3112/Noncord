package rabbitmq

import (
	"backend/internal/application/ports"
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RMQEventBroker struct {
	conn *amqp.Connection
}

func NewRMQEventBroker(conn *amqp.Connection) ports.EventsBroker {
	return &RMQEventBroker{conn}
}

func (mq *RMQEventBroker) Publish(ctx context.Context, msg ports.EventMessage) error {
	return fmt.Errorf("not implemented")
}

func (mq *RMQEventBroker) Close(ctx context.Context) error {
	return fmt.Errorf("not implemented")
}

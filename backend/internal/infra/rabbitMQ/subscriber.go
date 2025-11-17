package rabbitmq

import (
	"backend/internal/application/ports"
	_ "backend/internal/domain/entities"
	"backend/internal/domain/events"
	"context"
	"log/slog"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Handler func(context.Context, ports.EventMessage) error

type RMQEventSubscriber struct {
	mu           *sync.RWMutex
	c            *amqp.Channel
	logger       *slog.Logger
	serviceName  string
	exchangeName string
	handlerMap   map[string]Handler
	durable      bool
}

func NewRMQEventConsumer(ctx context.Context, conn *amqp.Connection, logger *slog.Logger, serviceName, exchangeName string, durable bool) (ports.EventSubscriber, error) {
	c, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	if err = c.ExchangeDeclare(exchangeName, amqp.ExchangeTopic, true, false, false, false, nil); err != nil {
		return nil, err
	}

	q, err := c.QueueDeclare(serviceName, durable, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	msgs, err := c.ConsumeWithContext(ctx, q.Name, "", false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	var mu sync.RWMutex
	client := &RMQEventSubscriber{&mu, c, logger, serviceName, exchangeName, make(map[string]Handler), durable}

	go client.consumeLoop(ctx, msgs)

	return client, nil
}

func (s *RMQEventSubscriber) Subscribe(ctx context.Context, topic string, handler func(context.Context, ports.EventMessage) error) error {
	q, err := s.c.QueueDeclarePassive(s.serviceName, s.durable, false, false, false, nil)
	if err != nil {
		return err
	}

	if err = s.c.QueueBind(q.Name, topic, s.exchangeName, false, nil); err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlerMap[topic] = handler
	return nil
}

func (s *RMQEventSubscriber) Close(ctx context.Context) error {
	return s.c.Close()
}

func (s *RMQEventSubscriber) consumeLoop(ctx context.Context, msgs <-chan amqp.Delivery) {
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case msg, ok := <-msgs:
			if !ok {
				slog.Info("Message channel closed")
				return
			}

			data := msg.Body
			base, err := events.ParseEvent(data)
			if err != nil {
				s.logger.Warn("Parse event failed", "error", err, "data", data)
				continue
			}

			s.mu.RLock()
			handler, ok := s.handlerMap[base.EventType]
			s.mu.RUnlock()
			if !ok {
				s.logger.Warn("Event don't have handler", "topic", base.EventType, "event", base)
				continue
			}

			err = handler(ctx, ports.EventMessage{
				AggregateId: base.AggregateID,
				EventType:   base.EventType,
				Payload:     data,
				Headers:     msg.Headers,
			})
			if err != nil {
				s.logger.Warn("Event handler failed", "event", base)
				continue
			}
		}
	}
}

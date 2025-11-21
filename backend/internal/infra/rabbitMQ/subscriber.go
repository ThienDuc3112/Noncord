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
	serviceName  string
	exchangeName string
	handlerMap   map[string]Handler
	durable      bool
}

func NewRMQEventSubscriber(ctx context.Context, conn *amqp.Connection, serviceName, exchangeName string, durable bool) (ports.EventSubscriber, error) {
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
	client := &RMQEventSubscriber{&mu, c, serviceName, exchangeName, make(map[string]Handler), durable}

	go client.consumeLoop(ctx, msgs)

	slog.Info("Created new EventSubscriber successfully")
	return client, nil
}

func (s *RMQEventSubscriber) Subscribe(topic string, handler func(context.Context, ports.EventMessage) error) error {
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

func (s *RMQEventSubscriber) Close() error {
	for topic := range s.handlerMap {
		s.c.QueueUnbind(s.serviceName, topic, s.exchangeName, nil)
	}

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
				slog.Warn("Parse event failed", "error", err, "data", data)
				msg.Ack(false)
				continue
			}

			s.mu.RLock()
			handler, ok := s.handlerMap[base.EventType]
			s.mu.RUnlock()
			if !ok {
				slog.Warn("Event don't have handler", "topic", base.EventType, "event", base)
				msg.Ack(false)
				continue
			}

			err = handler(ctx, ports.EventMessage{
				AggregateId: base.AggregateID,
				EventType:   base.EventType,
				Payload:     data,
				Headers:     msg.Headers,
			})
			if err != nil {
				slog.Warn("Event handler failed", "event", base)
				msg.Reject(true)
				continue
			} else {
				msg.Ack(false)
			}
		}
	}
}

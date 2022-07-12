package consumer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/thefuga/go-poc/internal/order/channel"
	"github.com/thefuga/go-poc/internal/order/event"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

// Consumer is a wrapper for an order consumer. Each order event will be associated
// with one consumer. See Consumer.RunConsumer.
type Consumer[T event.Event] struct {
	consumer *kafka.Consumer
}

func NewConsumer[T event.Event](config *kafka.ConfigMap) (*Consumer[T], error) {
	consumer, err := kafka.NewConsumer(config)

	if err != nil {
		return nil, err
	}

	return &Consumer[T]{
		consumer: consumer,
	}, nil
}

// SubscribeTopics is a convenience to help registering topics on each event consumer.
// See consumer_module.go.
func SubscribeTopics(consumer *kafka.Consumer, topics []string) error {
	return consumer.SubscribeTopics(topics, nil)
}

// RunConsumer sets the consumer associated with the specified event type to run
// on the application's startup.
func RunConsumer[T event.Event](
	c *Consumer[T], orderChan channel.OrderEventChannel[T], lifecycle fx.Lifecycle,
) {
	lifecycle.Append(fx.Hook{OnStart: func(context.Context) error {
		go c.Consume(orderChan)
		return nil
	}})
}

// Consume consumes the topic related to the consumer of the event type, and
// sends - after proper validation and error handling - the data to the channel
// associated with the event type, effectively decoupling the kafka consumer
// from event processing.
// See the package github.com/thefuga/go-poc/internal/order/processor for details
// on the actual event processing and github.com/thefuga/go-poc/internal/order/channel/channel_module.go
// for details on all registered channels/
func (c Consumer[T]) Consume(eventChan channel.OrderEventChannel[T]) {
	timeout := viper.GetDuration("app.kafka.consumers.order.timeout") * time.Millisecond

	for {
		message, consumerErr := c.consumer.ReadMessage(timeout)

		if consumerErr != nil {
			continue
		}

		var event T

		if unmarshalErr := json.Unmarshal(message.Value, &event); unmarshalErr != nil {
			continue
		}

		if validationErr := event.Validate(); validationErr != nil {
			continue
		}

		eventChan <- event
	}
}

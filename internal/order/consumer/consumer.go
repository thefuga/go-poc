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

type Consumer[T event.Event] struct {
	consumer *kafka.Consumer
}

func NewConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("app.kafka.address"),
		"group.id":          "1",
		"auto.offset.reset": "earliest",
	}
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

func SubscribeTopics(consumer *kafka.Consumer, topics []string) error {
	return consumer.SubscribeTopics(topics, nil)
}

func RunConsumer[T event.Event](
	c *Consumer[T], orderChan channel.OrderEventChannel[T], lifecycle fx.Lifecycle,
) {
	lifecycle.Append(fx.Hook{OnStart: func(context.Context) error {
		go c.Consume(orderChan)
		return nil
	}})
}

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

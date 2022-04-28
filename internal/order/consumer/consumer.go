package consumer

import (
	"context"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func NewConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("app.kafka.address"),
		"group.id":          "1",
		"auto.offset.reset": "earliest",
	}
}

func NewConsumer(config *kafka.ConfigMap) (*kafka.Consumer, error) {
	return kafka.NewConsumer(config)
}

func SubscribeTopics(consumer *kafka.Consumer, topics []string) error {
	return consumer.SubscribeTopics(topics, nil)
}

func RunConsumer(consumer *kafka.Consumer, lifecycle fx.Lifecycle) {
	lifecycle.Append(fx.Hook{OnStart: func(context.Context) error {
		go consume(consumer)
		return nil
	}})
}

func consume(consumer *kafka.Consumer) {
	timeout := viper.GetDuration("app.kafka.consumers.order.timeout") * time.Millisecond

	for {
		message, err := consumer.ReadMessage(timeout)

		if err != nil {
			continue
		}

		fmt.Printf(
			"Consumed event from topic %s: key = %-10s value = %s\n",
			*message.TopicPartition.Topic,
			string(message.Key),
			string(message.Value),
		)
	}
}

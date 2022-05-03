package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	func() *kafka.ConfigMap {
		return &kafka.ConfigMap{
			"bootstrap.servers": kafkaAddress(),
			"group.id":          "1",
			"auto.offset.reset": "earliest",
		}
	},
)

func kafkaAddress() string {
	address := viper.GetString("app.kafka.address")

	if address == "" {
		return "localhost"
	}

	return address
}

package producer

import (
	"context"
	"fmt"

	"github.com/thefuga/go-poc/internal/order/channel"
	"github.com/thefuga/go-poc/internal/order/event"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.uber.org/fx"
)

// Producer is a generic wrapper for the kafka.Producer. It allows the RunProducer
// method to be bind to an specific event type.
type Producer[T event.Event] struct {
	producer *kafka.Producer
}

func NewProducer[T event.Event](config *kafka.ConfigMap) (*Producer[T], error) {
	producer, err := kafka.NewProducer(config)

	if err != nil {
		return nil, err
	}

	return &Producer[T]{producer: producer}, nil
}

func (producer Producer[T]) Report() {
	go func() {
		for e := range producer.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
}

// RunProducer hooks the listening of the producer event channel to the application's
// lifecycle. In practice, it produces a kafka message from events received by the
// eventChan, on the given topic.
func (producer Producer[T]) RunProducer(
	eventChan channel.OrderEventChannel[T], topic string, lifecycle fx.Lifecycle,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				eventChan.Listen(func(event T) {
					bytes, _ := event.Bytes()

					message := &kafka.Message{
						Value: bytes,
						TopicPartition: kafka.TopicPartition{
							Topic: &topic, Partition: kafka.PartitionAny,
						},
					}

					fmt.Println(producer.producer.Produce(message, nil))
				})
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			producer.producer.Close()
			return nil
		},
	})
}

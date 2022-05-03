package producer_consumer_test

import (
	"context"
	"testing"
	"time"

	"github.com/thefuga/go-poc/configs"
	"github.com/thefuga/go-poc/internal/kafka"
	"github.com/thefuga/go-poc/internal/order/channel"
	"github.com/thefuga/go-poc/internal/order/consumer"
	"github.com/thefuga/go-poc/internal/order/event"
	"github.com/thefuga/go-poc/internal/order/producer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"go.uber.org/fx"
)

var (
	ProducerChans struct {
		fx.In

		Creation     channel.OrderEventChannel[event.Create] `name:"producer-creation"`
		Payment      channel.OrderEventChannel[event.Pay]    `name:"producer-payment"`
		Cancellation channel.OrderEventChannel[event.Cancel] `name:"producer-cancellation"`
	}
	ConsumerChans struct {
		fx.In

		Creation     channel.OrderEventChannel[event.Create] `name:"consumer-creation"`
		Payment      channel.OrderEventChannel[event.Pay]    `name:"consumer-payment"`
		Cancellation channel.OrderEventChannel[event.Cancel] `name:"consumer-cancellation"`
	}

	TestApp = fx.New(
		configs.Module,
		kafka.Module,
		producer.Module,
		consumer.Module,
		channel.Module,

		configs.Invokables,
		producer.Invokables,
		consumer.Invokables,

		fx.Populate(&ProducerChans),
		fx.Populate(&ConsumerChans),
	)
)

func TestOrderEventProducerConsumer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Order events producer-consumer suite")
}

var _ = BeforeSuite(func() {
	Expect(TestApp.Start(context.Background())).To(Succeed())
})

var _ = AfterSuite(func() {
	Expect(TestApp.Stop(context.Background())).To(Succeed())
})

var _ = Describe("order events producer-consumer", func() {
	Describe("sending creation events to the producer channel", func() {
		It("results in creation consumer channel receiveing the event", func() {
			ProducerChans.Creation <- event.Create{}

			Eventually(ConsumerChans.Creation).
				WithTimeout(100 * time.Second).
				Should(Receive())
		})
	})

	Describe("sending payment events to the producer channel", func() {
		It("results in payment consumer channel receiveing the event", func() {
			ProducerChans.Payment <- event.Pay{}

			Eventually(ConsumerChans.Payment).
				WithTimeout(100 * time.Second).
				Should(Receive())
		})
	})

	Describe("sending cancellation events to the producer channel", func() {
		It("results in cancellation consumer channel receiveing the event", func() {
			ProducerChans.Cancellation <- event.Cancel{}

			Eventually(ConsumerChans.Cancellation).
				WithTimeout(100 * time.Second).Should(Receive())
		})
	})
})

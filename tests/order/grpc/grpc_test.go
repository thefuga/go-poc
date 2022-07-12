package grpc_test

import (
	"context"
	_ "embed"
	"testing"

	"github.com/thefuga/go-poc/configs"
	grpcServer "github.com/thefuga/go-poc/internal/grpc"
	"github.com/thefuga/go-poc/internal/order/channel"
	"github.com/thefuga/go-poc/internal/order/event"
	orderHandler "github.com/thefuga/go-poc/internal/order/grpc"

	"github.com/spf13/viper"
	orderGRPC "go.buf.build/grpc/go/thefuga/go-poc/order/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"go.uber.org/fx"
)

var (
	Config *viper.Viper

	ProducerChans struct {
		fx.In

		Creation     channel.OrderEventChannel[event.Create] `name:"producer-creation"`
		Payment      channel.OrderEventChannel[event.Pay]    `name:"producer-payment"`
		Cancellation channel.OrderEventChannel[event.Cancel] `name:"producer-cancellation"`
	}

	TestApp = fx.New(
		fx.Provide(func() (*viper.Viper, error) {
			Config = viper.New()

			Config.Set("app.grpc_server.address", "localhost:3000")

			return Config, nil
		}),

		grpcServer.Module,
		orderHandler.Module,
		channel.Module,

		configs.Invokables,
		grpcServer.Invokables,
		orderHandler.Invokables,

		fx.Populate(&ProducerChans),
	)

	OrderClient orderGRPC.OrderServiceClient
	GRPCConn    *grpc.ClientConn
)

func TestOrderGRPCHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Order GRPC handler suite")
}

var _ = BeforeSuite(func() {
	Expect(TestApp.Start(context.Background())).To(Succeed())

	var dialErr error
	GRPCConn, dialErr = grpc.Dial(
		Config.GetString("app.grpc_server.address"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	Expect(dialErr).NotTo(HaveOccurred())

	OrderClient = orderGRPC.NewOrderServiceClient(GRPCConn)
})

var _ = AfterSuite(func() {
	Expect(GRPCConn.Close()).To(Succeed())
	Expect(TestApp.Stop(context.Background())).To(Succeed())
})

var _ = Describe("grpc calls on order handler", func() {
	Describe("Create method", func() {
		It("sends a create event to the creation producer channel", func() {
			var (
				respChan = make(chan *orderGRPC.CreateResponse, 1)
				errChan  = make(chan error)
			)

			go func() {
				resp, err := OrderClient.Create(
					context.Background(),
					&orderGRPC.CreateRequest{UserId: 1, ProductId: 2},
				)

				respChan <- resp
				errChan <- err
			}()

			Eventually(ProducerChans.Creation).Should(
				Receive(BeEquivalentTo(event.Create{UserID: 1, ProductID: 2})),
			)
			Expect(<-respChan).NotTo(BeNil())
			Expect(<-errChan).NotTo(HaveOccurred())
		})
	})
})

package grpc

import (
	"context"
	"net"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var (
	Module = fx.Provide(
		NewServer,
	)

	Invokables = fx.Invoke(InvokeServer)
)

func NewServer() *grpc.Server {
	return grpc.NewServer()
}

func InvokeServer(server *grpc.Server, lifecycle fx.Lifecycle, configs *viper.Viper) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			return startServer(server, configs.GetString("app.grpc_server.address"))
		},
		OnStop: func(context.Context) error {
			server.GracefulStop()
			return nil
		},
	})
}

func startServer(server *grpc.Server, address string) error {
	lis, err := net.Listen("tcp", address)

	if err != nil {
		return err
	}

	errChan := make(chan error)

	go func() {
		errChan <- server.Serve(lis)
	}()

	select {
	case <-time.After(time.Second):
		return nil
	case err := <-errChan:
		return err
	}
}

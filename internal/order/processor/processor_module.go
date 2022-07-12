/*
processor package has the definition for order events processors. These processors
won't necessairly contain business logic, but can (and should) be composed by the
services that actually perform business logic. I.e. these will just receive and
map events to the appropriate service without coupling any handler (be it gRPC, kafka, etc...)
to the domain.
*/
package processor

import (
	"github.com/thefuga/go-poc/internal/order/channel"

	"go.uber.org/fx"
)

// Module annotates processors with their respective consumer channels.
var Module = fx.Provide(
	fx.Annotate(
		NewCreationProcessor,
		fx.ParamTags(channel.CreationConsumer.Tag()),
	),

	fx.Annotate(
		NewPaymentProcessor,
		fx.ParamTags(channel.PaymentConsumer.Tag()),
	),

	fx.Annotate(
		NewCancellationProcessor,
		fx.ParamTags(channel.CancellationConsumer.Tag()),
	),
)

var Invokables = fx.Invoke(
	ProcessCreationEvents,
	ProcessPaymentEvents,
	ProcessCancellationEvents,
)

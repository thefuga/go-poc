package processor

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewCreationProcessor,
		fx.ParamTags(`name:"consumer-creation"`),
	),

	fx.Annotate(
		NewPaymentProcessor,
		fx.ParamTags(`name:"consumer-payment"`),
	),

	fx.Annotate(
		NewCancellationProcessor,
		fx.ParamTags(`name:"consumer-cancellation"`),
	),
)

var Invokables = fx.Invoke(
	ProcessCreationEvents,
	ProcessPaymentEvents,
	ProcessCancellationEvents,
)

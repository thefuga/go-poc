package processor

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewCreationProcessor,
		fx.ParamTags(`name:"creation"`),
	),

	fx.Annotate(
		NewPaymentProcessor,
		fx.ParamTags(`name:"payment"`),
	),

	fx.Annotate(
		NewCancellationProcessor,
		fx.ParamTags(`name:"cancellation"`),
	),
)

var Invokables = fx.Invoke(
	ProcessCreationEvents,
	ProcessPaymentEvents,
	ProcessCancellationEvents,
)

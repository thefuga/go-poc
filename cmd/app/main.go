package main

import (
	"github.com/thefuga/go-template/internal"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		internal.ApplicationModule,
		internal.ApplicationInvokables,
	).Run()
}

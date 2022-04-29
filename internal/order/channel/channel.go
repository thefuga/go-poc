package channel

import "github.com/thefuga/go-poc/internal/order/event"

type OrderEventChannel[T event.Event] chan T

func NewOrderChannel[T event.Event]() OrderEventChannel[T] {
	return make(chan T)
}

func (c OrderEventChannel[T]) Listen(react func(event T)) {
	for {
		select {
		case event := <-c:
			react(event)
		default:
		}
	}
}

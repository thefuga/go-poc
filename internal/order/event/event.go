package event

type Event interface {
	Create | Pay | Cancel

	Validate() error
	Bytes() ([]byte, error)
}

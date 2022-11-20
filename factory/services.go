package factory

type Services interface {
	Produce(devui string)
	Consume()
}

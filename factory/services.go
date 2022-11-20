package factory

type Services interface {
	Produce(devui string)
	Consume()
	BatchOf100() (devuis []string)
	ProduceBatch100(devuis []string)
}

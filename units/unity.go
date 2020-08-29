package units

type Unity struct {
	Value  float64
	Symbol string
}

func New(value float64, symbol string) Unity {
	return Unity{
		Value:  value,
		Symbol: symbol,
	}
}

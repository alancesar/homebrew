package bitterness

type Bitterness struct {
	IBU float64
}

func NewFromIBU(value float64) Bitterness {
	return Bitterness{IBU: value}
}

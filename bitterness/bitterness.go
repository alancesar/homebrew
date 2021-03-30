package bitterness

type Table map[string]Bitterness

type Bitterness struct {
	IBU float64
}

func (b Bitterness) IsZero() bool {
	return b.IBU == 0
}

func NewFromIBU(value float64) Bitterness {
	return Bitterness{IBU: value}
}

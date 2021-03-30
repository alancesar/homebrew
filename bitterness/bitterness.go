package bitterness

type (
	Table  []TableItem
	Method string

	TableItem struct {
		Method Method
		Value  Bitterness
	}

	Bitterness struct {
		IBU float64
	}
)

const (
	TinsethMethod Method = "Tinseth"
	RagerMethod   Method = "Rager"
	DanielMethod  Method = "Daniel"
)

func (b Bitterness) IsZero() bool {
	return b.IBU == 0
}

func NewFromIBU(value float64) Bitterness {
	return Bitterness{IBU: value}
}

package units

type Unity struct {
	Value  float64
	Symbol string
}

const (
	milligramSymbol = "mg"
	gramSymbol      = "g"
	kilogramSymbol  = "kg"
	poundSymbol     = "lb"
	ounceSymbol     = "oz"

	milliliterSymbol = "ml"
	literSymbol      = "l"
	gallonSymbol     = "gal"
)

func (Unity) Create(value float64, symbol string) Unity {
	return Unity{
		Value:  value,
		Symbol: symbol,
	}
}

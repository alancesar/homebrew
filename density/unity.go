package density

import "github.com/alancesar/homebrew/units"

const (
	sgSymbol   = "sg"
	brixSymbol = "ºbx"
)

type Density struct {
	Sg   units.Unity
	Brix units.Unity
}

func create(from *from) Density {
	return Density{
		Sg:   units.New(from.ToSg(), sgSymbol),
		Brix: units.New(from.ToBrix(), brixSymbol),
	}
}

func Sg(value float64) Density {
	return create(FromSg(value))
}

func Brix(value float64) Density {
	return create(FromBrix(value))
}

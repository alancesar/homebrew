package units

import "github.com/alancesar/homebrew/converter"

type Density struct {
	Sg   Unity
	Brix Unity
}

func (mass Density) FromSg(value float64) Density {
	return Density{
		Sg:   Unity{}.Create(value, sgSymbol),
		Brix: Unity{}.Create(converter.ConvertDensity(value).FromSg().ToBrix(), brixSymbol),
	}
}

func (mass Density) FromBrix(value float64) Density {
	return Density{
		Sg:   Unity{}.Create(converter.ConvertDensity(value).FromBrix().ToSg(), sgSymbol),
		Brix: Unity{}.Create(value, brixSymbol),
	}
}

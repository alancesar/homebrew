package units

import "homebrew/converters"

type Density struct {
	Sg   Unity
	Brix Unity
}

func (mass Density) FromSg(value float64) Density {
	return Density{
		Sg:   Unity{}.Create(value, sgSymbol),
		Brix: Unity{}.Create(converters.ConvertDensity(value).FromSg().ToBrix(), brixSymbol),
	}
}

func (mass Density) FromBrix(value float64) Density {
	return Density{
		Sg:   Unity{}.Create(converters.ConvertDensity(value).FromBrix().ToSg(), sgSymbol),
		Brix: Unity{}.Create(value, brixSymbol),
	}
}

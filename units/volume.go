package units

import "homebrew/calculator"

type Volume struct {
	Milliliters Unity
	Liters      Unity
	Gallons     Unity
}

func (volume Volume) FromMilliliter(value float64) Volume {
	return Volume{
		Milliliters: Unity{}.Create(value, milliliterSymbol),
		Liters:      Unity{}.Create(calculator.ConvertVolume(value).FromMilliliter().ToLiter(), literSymbol),
		Gallons:     Unity{}.Create(calculator.ConvertVolume(value).FromMilliliter().ToGallon(), gallonSymbol),
	}
}

func (volume Volume) FromLiter(value float64) Volume {
	return Volume{
		Milliliters: Unity{}.Create(calculator.ConvertVolume(value).FromLiter().ToMilliliter(), milliliterSymbol),
		Liters:      Unity{}.Create(value, literSymbol),
		Gallons:     Unity{}.Create(calculator.ConvertVolume(value).FromLiter().ToGallon(), gallonSymbol),
	}
}

func (volume Volume) FromGallon(value float64) Volume {
	return Volume{
		Milliliters: Unity{}.Create(calculator.ConvertVolume(value).FromGallon().ToMilliliter(), milliliterSymbol),
		Liters:      Unity{}.Create(calculator.ConvertVolume(value).FromGallon().ToLiter(), literSymbol),
		Gallons:     Unity{}.Create(value, gallonSymbol),
	}
}

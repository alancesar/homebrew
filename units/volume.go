package units

import "github.com/alancesar/homebrew/converter"

type Volume struct {
	Milliliters Unity
	Liters      Unity
	Gallons     Unity
}

func (volume Volume) FromMilliliter(value float64) Volume {
	return Volume{
		Milliliters: Unity{}.Create(value, milliliterSymbol),
		Liters:      Unity{}.Create(converter.ConvertVolume(value).FromMilliliter().ToLiter(), literSymbol),
		Gallons:     Unity{}.Create(converter.ConvertVolume(value).FromMilliliter().ToGallon(), gallonSymbol),
	}
}

func (volume Volume) FromLiter(value float64) Volume {
	return Volume{
		Milliliters: Unity{}.Create(converter.ConvertVolume(value).FromLiter().ToMilliliter(), milliliterSymbol),
		Liters:      Unity{}.Create(value, literSymbol),
		Gallons:     Unity{}.Create(converter.ConvertVolume(value).FromLiter().ToGallon(), gallonSymbol),
	}
}

func (volume Volume) FromGallon(value float64) Volume {
	return Volume{
		Milliliters: Unity{}.Create(converter.ConvertVolume(value).FromGallon().ToMilliliter(), milliliterSymbol),
		Liters:      Unity{}.Create(converter.ConvertVolume(value).FromGallon().ToLiter(), literSymbol),
		Gallons:     Unity{}.Create(value, gallonSymbol),
	}
}

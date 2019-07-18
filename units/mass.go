package units

import (
	"homebrew/calculator"
)

type Mass struct {
	Milligrams Unity
	Grams      Unity
	Kilograms  Unity
	Pounds     Unity
	Ounces     Unity
}

func (mass Mass) FromMilligram(value float64) Mass {
	return Mass{
		Milligrams: Unity{}.Create(value, milligramSymbol),
		Grams:      Unity{}.Create(calculator.ConvertMass(value).FromMilligram().ToGram(), gramSymbol),
		Kilograms:  Unity{}.Create(calculator.ConvertMass(value).FromMilligram().ToKilogram(), kilogramSymbol),
		Pounds:     Unity{}.Create(calculator.ConvertMass(value).FromMilligram().ToPounds(), poundSymbol),
		Ounces:     Unity{}.Create(calculator.ConvertMass(value).FromMilligram().ToOunces(), ounceSymbol),
	}
}

func (mass Mass) FromGram(value float64) Mass {
	return Mass{
		Milligrams: Unity{}.Create(calculator.ConvertMass(value).FromGram().ToMilligram(), milligramSymbol),
		Grams:      Unity{}.Create(value, gramSymbol),
		Kilograms:  Unity{}.Create(calculator.ConvertMass(value).FromGram().ToKilogram(), kilogramSymbol),
		Pounds:     Unity{}.Create(calculator.ConvertMass(value).FromGram().ToPounds(), poundSymbol),
		Ounces:     Unity{}.Create(calculator.ConvertMass(value).FromGram().ToOunces(), ounceSymbol),
	}
}

func (mass Mass) FromKilograms(value float64) Mass {
	return Mass{
		Milligrams: Unity{}.Create(calculator.ConvertMass(value).FromKilogram().ToMilligram(), milligramSymbol),
		Grams:      Unity{}.Create(calculator.ConvertMass(value).FromKilogram().ToGram(), gramSymbol),
		Kilograms:  Unity{}.Create(value, kilogramSymbol),
		Pounds:     Unity{}.Create(calculator.ConvertMass(value).FromKilogram().ToPounds(), poundSymbol),
		Ounces:     Unity{}.Create(calculator.ConvertMass(value).FromKilogram().ToOunces(), ounceSymbol),
	}
}

func (mass Mass) FromPounds(value float64) Mass {
	return Mass{
		Milligrams: Unity{}.Create(calculator.ConvertMass(value).FromPounds().ToMilligram(), milligramSymbol),
		Grams:      Unity{}.Create(calculator.ConvertMass(value).FromPounds().ToGram(), gramSymbol),
		Kilograms:  Unity{}.Create(calculator.ConvertMass(value).FromPounds().ToKilogram(), kilogramSymbol),
		Pounds:     Unity{}.Create(value, poundSymbol),
		Ounces:     Unity{}.Create(calculator.ConvertMass(value).FromPounds().ToOunces(), ounceSymbol),
	}
}

func (mass Mass) FromOunces(value float64) Mass {
	return Mass{
		Milligrams: Unity{}.Create(calculator.ConvertMass(value).FromOunces().ToMilligram(), milligramSymbol),
		Grams:      Unity{}.Create(calculator.ConvertMass(value).FromOunces().ToGram(), gramSymbol),
		Kilograms:  Unity{}.Create(calculator.ConvertMass(value).FromOunces().ToKilogram(), kilogramSymbol),
		Pounds:     Unity{}.Create(calculator.ConvertMass(value).FromOunces().ToPounds(), poundSymbol),
		Ounces:     Unity{}.Create(value, ounceSymbol),
	}
}
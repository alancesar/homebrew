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
	response := Mass{
		Milligrams: milligram(value),
		Grams:      gram(calculator.ConvertMass(value).FromMilligram().ToGram()),
		Kilograms:  kilogram(calculator.ConvertMass(value).FromMilligram().ToKilogram()),
		Pounds:     pound(calculator.ConvertMass(value).FromMilligram().ToPounds()),
		Ounces:     ounce(calculator.ConvertMass(value).FromMilligram().ToOunces()),
	}

	return response
}

func (mass Mass) FromGram(value float64) Mass {
	response := Mass{
		Milligrams: milligram(calculator.ConvertMass(value).FromGram().ToMilligram()),
		Grams:      gram(value),
		Kilograms:  kilogram(calculator.ConvertMass(value).FromGram().ToKilogram()),
		Pounds:     pound(calculator.ConvertMass(value).FromGram().ToPounds()),
		Ounces:     ounce(calculator.ConvertMass(value).FromGram().ToOunces()),
	}

	return response
}

func (mass Mass) FromKilograms(value float64) Mass {
	response := Mass{
		Milligrams: milligram(calculator.ConvertMass(value).FromKilogram().ToMilligram()),
		Grams:      gram(calculator.ConvertMass(value).FromKilogram().ToGram()),
		Kilograms:  kilogram(value),
		Pounds:     pound(calculator.ConvertMass(value).FromKilogram().ToPounds()),
		Ounces:     ounce(calculator.ConvertMass(value).FromKilogram().ToOunces()),
	}

	return response
}

func (mass Mass) FromPounds(value float64) Mass {
	response := Mass{
		Milligrams: milligram(calculator.ConvertMass(value).FromPounds().ToMilligram()),
		Grams:      gram(calculator.ConvertMass(value).FromPounds().ToGram()),
		Kilograms:  kilogram(calculator.ConvertMass(value).FromPounds().ToKilogram()),
		Pounds:     pound(value),
		Ounces:     ounce(calculator.ConvertMass(value).FromPounds().ToOunces()),
	}

	return response
}

func (mass Mass) FromOunces(value float64) Mass {
	response := Mass{
		Milligrams: milligram(calculator.ConvertMass(value).FromOunces().ToMilligram()),
		Grams:      gram(calculator.ConvertMass(value).FromOunces().ToGram()),
		Kilograms:  kilogram(calculator.ConvertMass(value).FromOunces().ToKilogram()),
		Pounds:     pound(calculator.ConvertMass(value).FromOunces().ToPounds()),
		Ounces:     ounce(value),
	}

	return response
}

func milligram(value float64) Unity {
	return Unity{
		Value:  value,
		Symbol: "mg",
	}
}

func gram(value float64) Unity {
	return Unity{
		Value:  value,
		Symbol: "g",
	}
}

func kilogram(value float64) Unity {
	return Unity{
		Value:  value,
		Symbol: "kg",
	}
}

func pound(value float64) Unity {
	return Unity{
		Value:  value,
		Symbol: "lb",
	}
}

func ounce(value float64) Unity {
	return Unity{
		Value:  value,
		Symbol: "oz",
	}
}

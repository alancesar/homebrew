package converters

const (
	milligramsInGrams = 1000
	kilogramsInGrams  = 0.001
	poundsInGrams     = 453.592
	ouncesInGrams     = 28.3495
	poundsInOunces    = 16
)

type FromMass struct {
	input float64
}

func (from FromMass) FromMilligram() *ToMass {
	grams := from.input / milligramsInGrams

	return &ToMass{
		grams:  grams,
		pounds: grams / poundsInGrams,
	}
}

func (from FromMass) FromGram() *ToMass {
	grams := from.input

	return &ToMass{
		grams:  grams,
		pounds: from.input / poundsInGrams,
	}
}

func (from FromMass) FromKilogram() *ToMass {
	grams := from.input / kilogramsInGrams

	return &ToMass{
		grams:  grams,
		pounds: grams / poundsInGrams,
	}
}

func (from FromMass) FromPounds() *ToMass {
	grams := from.input * poundsInGrams

	return &ToMass{
		grams:  grams,
		pounds: from.input,
	}
}

func (from FromMass) FromOunces() *ToMass {
	grams := from.input * ouncesInGrams

	return &ToMass{
		grams:  grams,
		pounds: from.input / poundsInOunces,
	}
}

type ToMass struct {
	grams  float64
	pounds float64
}

func (to ToMass) ToMilligram() float64 {
	return to.grams * milligramsInGrams
}

func (to ToMass) ToGram() float64 {
	return to.grams
}

func (to ToMass) ToKilogram() float64 {
	return to.grams * kilogramsInGrams
}

func (to ToMass) ToPounds() float64 {
	return to.pounds
}

func (to ToMass) ToOunces() float64 {
	return to.pounds * poundsInOunces
}

func ConvertMass(value float64) *FromMass {
	return &FromMass{
		input: value,
	}
}

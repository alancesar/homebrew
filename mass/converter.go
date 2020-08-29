package mass

const (
	milligramsInGrams = 1000
	kilogramsInGrams  = 0.001
	poundsInGrams     = 453.592
	ouncesInGrams     = 28.3495
	poundsInOunces    = 16
)

type from struct {
	grams  float64
	pounds float64
}

func (f *from) ToMilligram() float64 {
	return f.grams * milligramsInGrams
}

func (f *from) ToGram() float64 {
	return f.grams
}

func (f *from) ToKilogram() float64 {
	return f.grams * kilogramsInGrams
}

func (f *from) ToPounds() float64 {
	return f.pounds
}

func (f from) ToOunces() float64 {
	return f.pounds * poundsInOunces
}

func FromMilligram(value float64) *from {
	grams := value / milligramsInGrams

	return &from{
		grams:  grams,
		pounds: grams / poundsInGrams,
	}
}

func FromGram(value float64) *from {
	return &from{
		grams:  value,
		pounds: value / poundsInGrams,
	}
}

func FromKilogram(value float64) *from {
	grams := value / kilogramsInGrams

	return &from{
		grams:  grams,
		pounds: grams / poundsInGrams,
	}
}

func FromPound(value float64) *from {
	grams := value * poundsInGrams

	return &from{
		grams:  grams,
		pounds: value,
	}
}

func FromOunce(value float64) *from {
	grams := value * ouncesInGrams

	return &from{
		grams:  grams,
		pounds: value / poundsInOunces,
	}
}

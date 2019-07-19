package converter

const (
	milligramsInGrams = 1000
	kilogramsInGrams  = 0.001
	poundsInGrams     = 453.592
	ouncesInGrams     = 28.3495
	poundsInOunces    = 16
)

var (
	grams  float64
	pounds float64
)

type FromMass struct{}

func (from FromMass) FromMilligram() *ToMass {
	grams = input / milligramsInGrams
	pounds = grams / poundsInGrams
	return &ToMass{}
}

func (from FromMass) FromGram() *ToMass {
	grams = input
	pounds = input / poundsInGrams
	return &ToMass{}
}

func (from FromMass) FromKilogram() *ToMass {
	grams = input / kilogramsInGrams
	pounds = grams / poundsInGrams
	return &ToMass{}
}

func (from FromMass) FromPounds() *ToMass {
	grams = input * poundsInGrams
	pounds = input
	return &ToMass{}
}

func (from FromMass) FromOunces() *ToMass {
	grams = input * ouncesInGrams
	pounds = input / poundsInOunces
	return &ToMass{}
}

type ToMass struct{}

func (to ToMass) ToMilligram() float64 {
	return grams * milligramsInGrams
}

func (to ToMass) ToGram() float64 {
	return grams
}

func (to ToMass) ToKilogram() float64 {
	return grams * kilogramsInGrams
}

func (to ToMass) ToPounds() float64 {
	return pounds
}

func (to ToMass) ToOunces() float64 {
	return pounds * poundsInOunces
}

func ConvertMass(value float64) *FromMass {
	input = value
	return &FromMass{}
}

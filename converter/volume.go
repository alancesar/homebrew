package converter

const (
	millilitersInLiters = 1000
	gallonsInLiters     = 0.26417205235815
)

var (
	liters float64
)

type FromVolume struct{}

func (from FromVolume) FromMilliliter() *ToVolume {
	liters = input / millilitersInLiters
	return &ToVolume{}
}

func (from FromVolume) FromLiter() *ToVolume {
	liters = input
	return &ToVolume{}
}

func (from FromVolume) FromGallon() *ToVolume {
	liters = input / gallonsInLiters
	return &ToVolume{}
}

type ToVolume struct{}

func (to ToVolume) ToMilliliter() float64 {
	return liters * millilitersInLiters
}

func (to ToVolume) ToLiter() float64 {
	return liters
}

func (to ToVolume) ToGallon() float64 {
	return liters * gallonsInLiters
}

func ConvertVolume(value float64) *FromVolume {
	input = value
	return &FromVolume{}
}

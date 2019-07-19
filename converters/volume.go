package converters

const (
	millilitersInLiters = 1000
	gallonsInLiters     = 0.26417205235815
)

type FromVolume struct {
	input float64
}

func (from FromVolume) FromMilliliter() *ToVolume {
	return &ToVolume{
		liters: from.input / millilitersInLiters,
	}
}

func (from FromVolume) FromLiter() *ToVolume {
	return &ToVolume{
		liters: from.input,
	}
}

func (from FromVolume) FromGallon() *ToVolume {
	return &ToVolume{
		liters: from.input / gallonsInLiters,
	}
}

type ToVolume struct {
	liters float64
}

func (to ToVolume) ToMilliliter() float64 {
	return to.liters * millilitersInLiters
}

func (to ToVolume) ToLiter() float64 {
	return to.liters
}

func (to ToVolume) ToGallon() float64 {
	return to.liters * gallonsInLiters
}

func ConvertVolume(value float64) *FromVolume {
	return &FromVolume{
		input: value,
	}
}

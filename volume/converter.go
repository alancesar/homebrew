package volume

const (
	millilitersInLiters = 1000
	gallonsInLiters     = 0.26417205235815
)

type from struct {
	liters float64
}

func (f from) ToMilliliter() float64 {
	return f.liters * millilitersInLiters
}

func (f from) ToLiter() float64 {
	return f.liters
}

func (f from) ToGallon() float64 {
	return f.liters * gallonsInLiters
}

func FromMilliliter(value float64) *from {
	return &from{
		liters: value / millilitersInLiters,
	}
}

func FromLiter(value float64) *from {
	return &from{
		liters: value,
	}
}

func FromGallon(value float64) *from {
	return &from{
		liters: value / gallonsInLiters,
	}
}

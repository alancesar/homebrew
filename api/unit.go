package api

type Density struct {
	SG   float64 `json:"sg"`
	Brix float64 `json:"brix"`
}

type Color struct {
	SRM      float64 `json:"srm"`
	EBC      float64 `json:"ebc"`
	Lovibond float64 `json:"lovibond"`
}

type Alcohol struct {
	ABV         float64 `json:"abv"`
	ABW         float64 `json:"abw"`
	Attenuation float64 `json:"attenuation"`
}

type BitternessValue struct {
	IBU int `json:"ibu"`
}

type Bitterness struct {
	Method string          `json:"method"`
	Value  BitternessValue `json:"value"`
}

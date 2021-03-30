package api

type Recipe struct {
	UUID          string        `json:"uuid"`
	Name          string        `json:"name"`
	Efficiency    int           `json:"efficiency"`
	Attenuation   int           `json:"attenuation"`
	OG            float64       `json:"og"`
	FG            float64       `json:"fg"`
	WortCollected string        `json:"wort_collected"`
	BatchSize     string        `json:"batch_size"`
	Hops          []Hop         `json:"hops"`
	Fermentable   []Fermentable `json:"fermentable"`
}

type Hop struct {
	Quantity   string  `json:"quantity"`
	BoilTime   int     `json:"boil_time"`
	AlphaAcids float64 `json:"alpha_acids"`
	Pellet     bool    `json:"pellet"`
	DryHopping bool    `json:"dry_hopping"`
}

type Fermentable struct {
	Quantity string  `json:"quantity"`
	Lovibond float64 `json:"lovibond"`
	PPG      float64 `json:"ppg"`
	Mashing  bool    `json:"mashing"`
}

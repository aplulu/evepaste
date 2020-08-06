package entity

type Price struct {
	TypeId int32 `json:"type_id"`
	Buy PriceItem `json:"buy"`
	All PriceItem `json:"all"`
	Sell PriceItem `json:"sell"`
}

type PriceItem struct {
	Avg float64 `json:"avg"`
	Median float64 `json:"median"`
	Max float64 `json:"max"`
	Min float64 `json:"min"`
	Price float64 `json:"price"`
}

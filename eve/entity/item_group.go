package entity

type ItemGroup struct {
	Key string `json:"key"`
	Name string `json:"name"`
	Items []Item `json:"items"`
	TotalBuyPrice float64 `json:"total_buy_price"`
	TotalSellPrice float64 `json:"total_sell_price"`
	TotalVolume float64 `json:"total_volume"`
}

func NewItemGroup(key string, name string) *ItemGroup {
	g := &ItemGroup{
		Key: key,
		Name: name,
	}
	g.Items = make([]Item, 0)

	return g
}
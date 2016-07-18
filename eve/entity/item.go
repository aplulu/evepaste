package entity

type Item struct {
	TypeId int32 `json:"type_id"`
	TypeName string `json:"type_name"`
	GroupId int32 `json:"group_id"`
	GroupName string `json:"group_name"`
	Quantity int `json:"quantity"`
	Volume float64 `json:"volume"`
	Prices Price `json:"-"`
	Name string `json:"name,omitempty"`
	Distance int64 `json:"distance,omitempty"`
	SlotType string `json:"slot_type,omitempty"`
	ChargeTypeId int32 `json:"charge_type_id,omitempty"`
}

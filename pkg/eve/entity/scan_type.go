package entity

type ScanType struct {
	TypeId int32 `json:"type_id"`
	TypeName string `json:"type_name"`
	GroupId int32 `json:"group_id"`
	CategoryId int32 `json:"category_id"`
	OnGridCount int `json:"ongrid_count"`
	OffGridCount int `json:"offgrid_count"`
	Total int `json:"total"`
}

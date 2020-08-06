package entity

type ScanGroup struct {
	GroupId int32 `json:"group_id"`
	GroupName string `json:"group_name"`
	CategoryId int32 `json:"category_id"`
	OnGridCount int `json:"ongrid_count"`
	OffGridCount int `json:"offgrid_count"`
	Total int `json:"total"`
}
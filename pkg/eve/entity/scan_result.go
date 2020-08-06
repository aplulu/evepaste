package entity

type ScanResult struct {
	Types []ScanType `json:"types"`
	Groups []ScanGroup `json:"groups"`
	OnGridCount int `json:"ongrid_count"`
	OffGridCount int `json:"offgrid_count"`
	Total int `json:"total"`
	SolarSystemName string `json:"solarsystem_name"`
	SolarSystemSecurity float64 `json:"solarsystem_security"`
	ConstellationName string `json:"constellation_name"`
	RegionName string `json:"region_name"`
}

package fit

import (
	"fareastdominions.com/evepaste/eve/entity"
	"encoding/json"
)

var clfSlots = []string{
	"low",
	"med",
	"high",
	"rig",
	"subsystem",
}

type CLF struct {
	Version int `json:"clf-version"`
	GeneratedBy string `json:"X-generatedby,omitempty"`
	Metadata CLFMeta `json:"metadata,omitempty"`
	Ship CLFModule `json:"ship"`
	Presets []CLFPreset `json:"presets,omitempty"`
	Drones []CLFDronePreset `json:"drones,omitempty"`
}

type CLFMeta struct {
	Title string `json:"title"`
}

type CLFPreset struct {
	Name string `json:"presetname"`
	Modules []CLFModule `json:"modules"`
}

type CLFModule struct {
	TypeId int32 `json:"typeid"`
	Charges []CLFModule `json:"charges,omitempty"`
	CpId int32 `json:"cpid,omitempty"`
	Quantity int `json:"quantity,omitempty"`
}

type CLFChargePreset struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type CLFDronePreset struct {
	Name string `json:"presetname"`
	InBay []CLFModule `json:"inbay,omitempty"`
	InSpace []CLFModule `json:"inspace,omitempty"`
}


func ExportCLF(f *entity.Fit) string {
	modules := make([]CLFModule, 0)
	for _, slot := range clfSlots {
		for _, item := range f.GetItems(slot) {
			for i := 0; i < item.Quantity; i++ {
				module := CLFModule{
					TypeId: item.TypeId,
				}

				if item.ChargeTypeId != 0 {
					module.Charges = []CLFModule{{TypeId: item.ChargeTypeId}}
				}

				modules = append(modules, module)
			}
		}
	}

	drones := make([]CLFModule, 0)
	for _, item := range f.GetItems("drone") {
		drones = append(drones, CLFModule{TypeId: item.TypeId, Quantity: item.Quantity})
	}

	clf := &CLF{
		Version: 1,
		GeneratedBy: "EVEPaste",
		Metadata: CLFMeta{
			Title: f.Ship.Name,
		},
		Ship: CLFModule{
			TypeId: f.Ship.TypeId,
		},
		Presets: []CLFPreset{{
			Name: "Default preset",
			Modules: modules,
		}},
		Drones: []CLFDronePreset{{
			Name: "Default drone preset",
			InBay: drones,
		}},
	}


	bytes, err := json.Marshal(clf)
	if err != nil {
		return ""
	} else {
		return string(bytes)
	}
}

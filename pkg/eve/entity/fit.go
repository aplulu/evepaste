package entity

type FitSlot struct {
	Name string
	Items []Item
}

type Fit struct {
	Ship Item `json:"ship"`
	High []Item
	Med []Item
	Low []Item
	Rig []Item
	Subsystem []Item
	Drone []Item
	Other []Item
	Slots []FitSlot
}

func (s *Fit) GetItems(slot string) []Item {
	switch (slot) {
	case "ship":
		return []Item{s.Ship}
	case "high":
		return s.High
	case "med":
		return s.Med
	case "low":
		return s.Low
	case "rig":
		return s.Rig
	case "subsystem":
		return s.Subsystem
	case "drone":
		return s.Drone
	case "other":
		return s.Other
	}
	return make([]Item, 0)
}

func NewFitFromItems(items []Item) *Fit {
	fit := &Fit{}

	for _, item := range items {
		switch (item.SlotType) {
		case "ship":
			fit.Ship = item
		case "high":
			fit.High = append(fit.High, item)
		case "med":
			fit.Med = append(fit.Med, item)
		case "low":
			fit.Low = append(fit.Low, item)
		case "rig":
			fit.Rig = append(fit.Rig, item)
		case "subsystem":
			fit.Subsystem = append(fit.Subsystem, item)
		case "drone":
			fit.Drone = append(fit.Drone, item)
		default:
			fit.Other = append(fit.Other, item)
		}
	}

	slotNames := []string{"ship", "high", "med", "low", "rig", "subsystem", "drone", "other"}
	for _, slotName := range slotNames {
		items := fit.GetItems(slotName)
		if len(items) > 0 {
			fit.Slots = append(fit.Slots, FitSlot{
				Name: slotName,
				Items: items,
			})
		}
	}

	return fit
}
package fit

import (
	"github.com/evepaste/evepaste/pkg/eve/entity"
	"github.com/evepaste/evepaste/pkg/eve/repository"
	"fmt"
)

var eftSlots = []string{
	"low",
	"med",
	"high",
	"rig",
	"subsystem",
}

func ExportEFT(f *entity.Fit, lang string) string {
	ship := repository.GetTypeByID(f.Ship.TypeId)
	shipName := "No Ship"

	if ship != nil {
		shipName = ship.GetName(lang)
	}

	fitName := f.Ship.Name
	if fitName == "" {
		fitName = "New Ship"
	}

	text := fmt.Sprintf("[%s, %s]\n", shipName, fitName)

	for _, slot := range eftSlots {
		for _, item := range f.GetItems(slot) {
			t := repository.GetTypeByID(item.TypeId)
			if t != nil {
				for i := 0; i < item.Quantity; i++ {
					charge := ""
					if item.ChargeTypeId != 0 {
						ct := repository.GetTypeByID(item.ChargeTypeId)
						if ct != nil {
							charge = ct.GetName(lang)
						}
					}

					if charge != "" {
						text += fmt.Sprintf("%s, %s\n", t.GetName(lang), charge)
					} else {
						text += fmt.Sprintf("%s\n", t.GetName(lang))
					}
				}
			}
		}
		text += "\n"
	}

	for _, item := range f.GetItems("drone") {
		t := repository.GetTypeByID(item.TypeId)
		if t != nil {
			text += fmt.Sprintf("%s x%d\n", t.GetName(lang), item.Quantity)
		}
	}

	return text
}

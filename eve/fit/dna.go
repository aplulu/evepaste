package fit

import (
	"fareastdominions.com/evepaste/eve/entity"
	"strconv"
	"fmt"
)

var dnaSlots = []string{
	"low",
	"med",
	"high",
	"rig",
	"subsystem",
	"drone",
}

func ExportDNA(f *entity.Fit) string {
	out := ""

	if f.Ship.TypeId > 0 {
		out += strconv.Itoa(int(f.Ship.TypeId))
	}

	moduleMap := make(map[int32]int)
	for _, slot := range dnaSlots {
		for _, item := range f.GetItems(slot) {
			if _, ok := moduleMap[item.TypeId]; ok {
				moduleMap[item.TypeId] += item.Quantity
			} else {
				moduleMap[item.TypeId] = item.Quantity
			}
		}
	}

	for id, q := range moduleMap {
		out += fmt.Sprintf(":%d;%d", int(id), q)
	}


	return out + "::"
}

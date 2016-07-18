package parser

import (
	"testing"
	"fareastdominions.com/evepaste/eve/test"
	"fareastdominions.com/evepaste/eve/entity"
)

func TestInventoryParse(t *testing.T) {
	testCases := []test.ParserTestCase{{
		`Hydrogen Fuel Block	13,911	Fuel Block			69,555 m3
Minmatar Republic Starbase Charter	1,226	Lease			122.60 m3`,
		[]entity.Item{
			entity.Item{
				TypeName: `Hydrogen Fuel Block`,
				Quantity: 13911,
			},
			entity.Item{
				TypeName: `Minmatar Republic Starbase Charter`,
				Quantity: 1226,
			},
		},
	}}

	p := &InventoryParser{}
	for _, c := range testCases {
		items, err := p.Parse(c.GetLines())
		if err != nil {
			t.Fatal(err)
		}
		c.Assert(t, items)
	}

}

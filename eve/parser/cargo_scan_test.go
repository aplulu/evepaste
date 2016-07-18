package parser

import (
	"testing"
	"fareastdominions.com/evepaste/eve/test"
	"fareastdominions.com/evepaste/eve/entity"
)

func TestCargoScanParse(t *testing.T) {
	testCases := []test.ParserTestCase{{
		`5000 Caldari Navy Mjolnir Rocket
47720 Azure Plagioclase
5 Navy Cap Booster 400`,
		[]entity.Item{
			entity.Item{
				TypeName: `Caldari Navy Mjolnir Rocket`,
				Quantity: 5000,
			},
			entity.Item{
				TypeName: `Azure Plagioclase`,
				Quantity: 47720,
			},
			entity.Item{
				TypeName: `Navy Cap Booster 400`,
				Quantity: 5,
			},
		},
	}}

	p := &CargoScanParser{}
	for _, c := range testCases {
		items, err := p.Parse(c.GetLines())
		if err != nil {
			t.Fatal(err)
		}
		c.Assert(t, items)
	}
}

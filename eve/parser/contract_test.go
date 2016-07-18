package parser

import (
	"testing"
	"fareastdominions.com/evepaste/eve/test"
	"fareastdominions.com/evepaste/eve/entity"
)

func TestContractParse(t *testing.T) {
	testCases := []test.ParserTestCase{{
		`Ship Maintenance Array	1	Ship Maintenance Array	Starbase
Standup ASML Missile Launcher I	2	Structure ASML Missile Launcher	Structure Module	`,
		[]entity.Item{
			entity.Item{
				TypeName: `Ship Maintenance Array`,
				Quantity: 1,
			},
			entity.Item{
				TypeName: `Standup ASML Missile Launcher I`,
				Quantity: 2,
			},
		},
	}}

	p := &ContractParser{}
	for _, c := range testCases {
		items, err := p.Parse(c.GetLines())
		if err != nil {
			t.Fatal(err)
		}
		c.Assert(t, items)
	}
}

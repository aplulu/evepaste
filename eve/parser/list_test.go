package parser

import (
	"testing"
	"fareastdominions.com/evepaste/eve/test"
	"fareastdominions.com/evepaste/eve/entity"
)

func TestListParse(t *testing.T) {
	testCases := []test.ParserTestCase{{
		`167 x Liquid Ozone
20 x Strontium Clathrates`,
		[]entity.Item{
			entity.Item{
				TypeName: `Liquid Ozone`,
				Quantity: 167,
			},
			entity.Item{
				TypeName: `Strontium Clathrates`,
				Quantity: 20,
			},
		},
	}}

	p := &ListParser{}
	for _, c := range testCases {
		items, err := p.Parse(c.GetLines())
		if err != nil {
			t.Fatal(err)
		}
		c.Assert(t, items)
	}

}

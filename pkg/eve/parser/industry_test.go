package parser

import (
	"testing"
	"github.com/evepaste/evepaste/eve/test"
	"github.com/evepaste/evepaste/eve/entity"
)

func TestIndustryParse(t *testing.T) {
	testCases := []test.ParserTestCase{{
		`Tritanium (2054 Units)
Pyerite (2088 Units)`,
		[]entity.Item{
			entity.Item{
				TypeName: `Tritanium`,
				Quantity: 2054,
			},
			entity.Item{
				TypeName: `Pyerite`,
				Quantity: 2088,
			},
		},
	}}

	p := &IndustryParser{}
	for _, c := range testCases {
		items, err := p.Parse(c.GetLines())
		if err != nil {
			t.Fatal(err)
		}
		c.Assert(t, items)
	}

}

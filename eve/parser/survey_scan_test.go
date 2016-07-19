package parser

import (
	"testing"
	"fareastdominions.com/evepaste/eve/test"
	"fareastdominions.com/evepaste/eve/entity"
)

func TestSurvayScan(t *testing.T) {
	testCases := []test.ParserTestCase{{
		`Scordite	39,856	25 km
Scordite	8,320	22 km
Silvery Omber	2,649	23 km`,
		[]entity.Item{
			entity.Item{
				TypeName: `Scordite`,
				Quantity: 39856,
			},
			entity.Item{
				TypeName: `Scordite`,
				Quantity: 8320,
			},
			entity.Item{
				TypeName: `Silvery Omber`,
				Quantity: 2649,
			},
		},
	}}

	p := &SurveyScanParser{}
	for _, c := range testCases {
		items, err := p.Parse(c.GetLines())
		if err != nil {
			t.Fatal(err)
		}
		c.Assert(t, items)
	}
}
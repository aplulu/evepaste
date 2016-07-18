package parser

import (
	"testing"
	"fareastdominions.com/evepaste/eve/test"
	"fareastdominions.com/evepaste/eve/entity"
)

func TestDscanParse(t *testing.T) {
	testCases := []test.ParserTestCase{{
		`Frosty	Bifrost	1,002 km
Best Squad	Caldari Control Tower Small	-`,
		[]entity.Item{
			entity.Item{
				TypeName: `Bifrost`,
				Quantity: 1,
			},
			entity.Item{
				TypeName: `Caldari Control Tower Small`,
				Quantity: 1,
			},
		},
	}}

	p := &DscanParser{}
	for _, c := range testCases {
		items, err := p.Parse(c.GetLines())
		if err != nil {
			t.Fatal(err)
		}
		c.Assert(t, items)
	}

}

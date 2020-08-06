package parser

import (
	"testing"
	"github.com/evepaste/evepaste/eve/test"
	"github.com/evepaste/evepaste/eve/entity"
)

func TestDscanParse(t *testing.T) {
	testCases := []test.ParserTestCase{{
		`14	Earled X - Moon 9	Moon	2,188,990 km
26561	Gistii Hijacker Wreck	Angel Small Wreck	-`,
		[]entity.Item{
			entity.Item{
				TypeName: `Moon`,
				Quantity: 1,
			},
			entity.Item{
				TypeName: `Angel Small Wreck`,
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

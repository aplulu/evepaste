package parser

import (
	"testing"
	"github.com/evepaste/evepaste/eve/test"
	"github.com/evepaste/evepaste/eve/entity"
)

func TestViewContentsParse(t *testing.T) {
	testCases := []test.ParserTestCase{{
		`Compressed Pure Jaspet	Jaspet	23
Compressed Silvery Omber	Omber	1688
Kernite	Kernite	23199`,
		[]entity.Item{
			entity.Item{
				TypeName: `Compressed Pure Jaspet`,
				Quantity: 23,
			},
			entity.Item{
				TypeName: `Compressed Silvery Omber`,
				Quantity: 1688,
			},
			entity.Item{
				TypeName: `Kernite`,
				Quantity: 23199,
			},
		},
	}}

	p := &ViewContentsParser{}
	for _, c := range testCases {
		items, err := p.Parse(c.GetLines())
		if err != nil {
			t.Fatal(err)
		}
		c.Assert(t, items)
	}
}

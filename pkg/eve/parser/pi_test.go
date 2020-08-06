package parser

import (
	"testing"
	"github.com/evepaste/evepaste/eve/entity"
	"github.com/evepaste/evepaste/eve/test"
)


func TestPIParse(t *testing.T) {
	testCases := []test.ParserTestCase{{
		`	Livestock	430.0	645.0
	Proteins	540.0	205.2`,
		[]entity.Item{
			entity.Item{
				TypeName: `Livestock`,
				Quantity: 430,
			},
			entity.Item{
				TypeName: `Proteins`,
				Quantity: 540,
			},
		},
	}}

	p := &PIParser{}
	for _, c := range testCases {
		items, err := p.Parse(c.GetLines())
		if err != nil {
			t.Fatal(err)
		}
		c.Assert(t, items)
	}
}
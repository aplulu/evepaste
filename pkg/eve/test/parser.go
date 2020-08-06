package test

import (
	"strings"
	"testing"
	"github.com/evepaste/evepaste/eve/entity"
)

type ParserTestCase struct {
	Text string
	Items []entity.Item
}

func (c *ParserTestCase) GetLines() ([]string) {
	lines := strings.Split(c.Text, "\n")
	for i, _ := range lines {
		lines[i] = strings.TrimRight(lines[i], "\r\n")
	}
	return lines
}

func (c *ParserTestCase) Assert(t *testing.T, items []entity.Item) {
	if len(items) != len(c.Items) {
		t.Errorf("Wrong result count. got=%d, want=%d", len(items), len(c.Items))
	}

	for i := 0; i < len(items); i++ {
		if items[i].TypeName != c.Items[i].TypeName {
			t.Errorf("Wrong type name. got=%s, want=%s", items[i].TypeName, c.Items[i].TypeName)
		}

		if items[i].Quantity != c.Items[i].Quantity {
			t.Errorf("Wrong quantity. got=%d, want=%d", items[i].Quantity, c.Items[i].Quantity)
		}
	}
}


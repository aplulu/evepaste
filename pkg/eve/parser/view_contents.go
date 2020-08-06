package parser

import (
	"regexp"
	"strings"
	"github.com/evepaste/evepaste/pkg/eve/entity"
)


var ViewContentsRegex *regexp.Regexp = regexp.MustCompile(`^([\w ]+)*?\t([\w ]+)*?\t(\d+)$`)

type ViewContentsParser struct {
	BaseParser
}

func (p *ViewContentsParser) Type() string {
	return "view_contents"
}

func (p *ViewContentsParser) Parse(lines []string) ([]entity.Item, error) {
	items := []entity.Item{}

	for _, line := range lines {
		results := ViewContentsRegex.FindAllStringSubmatch(strings.TrimSpace(line), -1)

		if len(results) == 1 {
			quantity, err := p.parseInt(results[0][3])
			if err != nil {
				continue
			}

			items = append(items, entity.Item{
				TypeName: results[0][1],
				Quantity: quantity,
			})
		}
	}

	return items, nil
}
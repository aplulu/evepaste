package parser

import (
	"regexp"
	"strings"
	"fareastdominions.com/evepaste/eve/entity"
)


/*
4 x Enriched Uranium
22 x Oxygen
4 x Mechanical Parts
9 x Coolant
1 x Robotics
167 x Heavy Water
167 x Liquid Ozone
20 x Strontium Clathrates
*/

var ListRegex *regexp.Regexp = regexp.MustCompile(`^([\d,]+)\sx\s([\S ]+)$`)

type ListParser struct {
	BaseParser
}

func (p *ListParser) Type() string {
	return "list"
}


func (p *ListParser) Parse(lines []string) ([]entity.Item, error) {
	items := []entity.Item{}

	for _, line := range lines {
		results := ListRegex.FindAllStringSubmatch(strings.TrimSpace(line), -1)

		if len(results) == 1 {
			quantity, err := p.parseInt(results[0][1])
			if err != nil {
				continue
			}

			items = append(items, entity.Item{
				TypeName: strings.TrimRight(results[0][2], "*"),
				Quantity: quantity,
			})
		}
	}

	return items, nil
}
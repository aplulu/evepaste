package parser

import (
	"regexp"
	"strings"
	"fareastdominions.com/evepaste/eve/entity"
)

/**
Fertilizer	1,585	Refined Commodities			2,377.50 m3
Synthetic Oil	2,730	Refined Commodities			4,095 m3
 */

var InventoryRegex *regexp.Regexp = regexp.MustCompile(`^([\S ]+)\t([\d,\.]+)\t([\S ]+)\t([\S ]*)\t([\S ]*)\t([\d,\.]+ m3)$`)

type InventoryParser struct {
	BaseParser
}

func (p *InventoryParser) Type() string {
	return "inventory"
}

func (p *InventoryParser) Parse(lines []string) ([]entity.Item, error) {
	items := []entity.Item{}

	for _, line := range lines {
		results := InventoryRegex.FindAllStringSubmatch(strings.TrimSpace(line), -1)

		if len(results) == 1 && len(results[0]) == 7 {
			quantity, err := p.parseInt(results[0][2])
			if err != nil {
				continue
			}


			items = append(items, entity.Item{
				TypeName: strings.TrimRight(results[0][1], "*"),
				Quantity: quantity,
			})
		}
	}

	return items, nil
}
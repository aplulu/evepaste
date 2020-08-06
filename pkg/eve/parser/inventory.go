package parser

import (
	"regexp"
	"strings"
	"github.com/evepaste/evepaste/pkg/eve/entity"
)

/**
Fertilizer	1,585	Refined Commodities			2,377.50 m3
Synthetic Oil	2,730	Refined Commodities			4,095 m3
Tungsten	6,000	Moon Materials			2,400 m3	25,430,640.00 ISK
X-Large Neutron Saturation Injector I	1	Shield Booster		Medium	50 m3	64,603.24 ISK
X5 Enduring Stasis Webifier	1	Stasis Web		Medium	5 m3	8,819.62 ISK
X5 Enduring Stasis Webifier	2	Stasis Web		Medium	10 m3	17,639.24 ISK
 */

var InventoryRegex *regexp.Regexp = regexp.MustCompile(`^([\S ]+)\t([\d,\.]+)\t([\S ]+)\t([\S ]*)\t([\S ]*)\t([\d,\.]+ m3)(\t[\d,\.]+ ISK)?$`)

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

		if len(results) == 1 && len(results[0]) == 8 {
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
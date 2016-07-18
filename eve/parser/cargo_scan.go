package parser

import (
	"regexp"
	"strings"
	"fareastdominions.com/evepaste/eve/entity"
)


var CargoScanRegex *regexp.Regexp = regexp.MustCompile(`^([\d,\.]+)\s([\S ]+)(\s\((Copy|Original| Копия|)\))?$`)

type CargoScanParser struct {
	BaseParser
}

func (p *CargoScanParser) Type() string {
	return "cargo_scan"
}

func (p *CargoScanParser) Parse(lines []string) ([]entity.Item, error) {
	items := []entity.Item{}

	for _, line := range lines {
		results := CargoScanRegex.FindAllStringSubmatch(strings.TrimSpace(line), -1)

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
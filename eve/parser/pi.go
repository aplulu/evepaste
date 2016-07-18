package parser

import (
	"regexp"
	"strings"
	"fareastdominions.com/evepaste/eve/entity"
)

/**
	Livestock	430.0	645.0
	Proteins	540.0	205.2
 */


var PIRegex *regexp.Regexp = regexp.MustCompile(`^\t([\S ]+)\t([\d\.]+)\t([\d\.]+)$`)

type PIParser struct {
	BaseParser
}

func (p *PIParser) Type() string {
	return "pi"
}

func (p *PIParser) Parse(lines []string) ([]entity.Item, error) {
	items := []entity.Item{}

	for _, line := range lines {
		results := PIRegex.FindAllStringSubmatch(line, -1)

		if len(results) == 1 {
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

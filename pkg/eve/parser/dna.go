package parser

import (
	"regexp"
	"strings"
	"github.com/evepaste/evepaste/pkg/eve/entity"
)


var DNARegex *regexp.Regexp = regexp.MustCompile(`^(\d+):([0-9:;]+):$`)


type DNAParser struct {
	BaseParser
}

func (p *DNAParser) Type() string {
	return "fit"
}

func (p *DNAParser) Parse(lines []string) ([]entity.Item, error) {
	items := []entity.Item{}

	for i, line := range lines {
		if i == 0 {
			results := DNARegex.FindAllStringSubmatch(line, -1)
			if len(results) != 1 {
				return items, nil
			}

			shipId, err := p.parseInt(results[0][1])
			if err != nil {
				return items, nil
			}

			items = append(items, entity.Item{
				TypeId: int32(shipId),
				Quantity: 1,
				SlotType: "ship",
			})

			parts := strings.SplitN(results[0][2], ":", -1)
			for _, part := range parts {
				pairs := strings.SplitN(part, ";", 2)
				id, err := p.parseInt(pairs[0])
				if err != nil {
					continue
				}

				quantity := 1
				if len(pairs) > 1 {
					q, err := p.parseInt(pairs[1])
					if err == nil {
						quantity = q
					}
				}


				items = append(items, entity.Item{
					TypeId: int32(id),
					Quantity: quantity,
				})
			}
		} else {
			break
		}
	}

	return items, nil
}
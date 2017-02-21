package parser

import (
	"regexp"
	"strings"
	"fareastdominions.com/evepaste/eve/entity"
)

var LootHistoryRegex *regexp.Regexp = regexp.MustCompile(`^(\d{2}:\d{2}:\d{2}) ([\S ]+?) (has looted|が|забрала) ([\d,]+) x ([\S ]+)( をルート)?$`)

type LootHistoryParser struct {
	BaseParser
}

func (p *LootHistoryParser) Type() string {
	return "loot_history"
}

func (p *LootHistoryParser) Parse(lines []string) ([]entity.Item, error) {
	items := []entity.Item{}

	for _, line := range lines {
		results := LootHistoryRegex.FindAllStringSubmatch(strings.TrimSpace(line), -1)

		if len(results) == 1 {
			quantity, err := p.parseInt(results[0][4])
			if err != nil {
				continue
			}

			items = append(items, entity.Item{
				TypeName: strings.TrimRight(results[0][5], "*"),
				Quantity: quantity,
				ItemGroupKey: results[0][2],
				ItemGroupName: results[0][2],
			})
		}
	}

	return items, nil
}
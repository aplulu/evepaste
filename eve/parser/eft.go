package parser

import (
	"regexp"
	"strings"
	"fareastdominions.com/evepaste/eve/entity"
	"fareastdominions.com/evepaste/eve/repository"
)


var EFTFirstRegex *regexp.Regexp = regexp.MustCompile(`^\[([\S ]+), ([\S ]+)\]$`)
var EFTItemRegex *regexp.Regexp = regexp.MustCompile(`^([\S ]+?)( x(\d+))?(, ([\S ]+))?$`)

type EFTParser struct {
	BaseParser
}

func (p *EFTParser) Type() string {
	return "fit"
}

func (p *EFTParser) Parse(lines []string) ([]entity.Item, error) {
	items := []entity.Item{}

	for i, line := range lines {
		if i == 0 {
			results := EFTFirstRegex.FindAllStringSubmatch(line, -1)
			if len(results) != 1 {
				return items, nil
			}
			items = append(items, entity.Item{
				TypeName: strings.TrimRight(results[0][1], "*"),
				Quantity: 1,
				Name: results[0][2],
				SlotType: "ship",
			})
		} else if line == "" {
			continue
		} else {
			results := EFTItemRegex.FindAllStringSubmatch(line, -1)
			if len(results) != 1 {
				continue
			}

			quantity := 1
			if results[0][3] != "" {
				q, err := p.parseInt(results[0][3])
				if err != nil {
					continue
				}
				quantity = q
			}

			item := entity.Item{
				TypeName: strings.TrimRight(results[0][1], "*"),
				Quantity: quantity,
			}

			// Charges
			if results[0][5] != "" {
				t := repository.GetTypeByName(strings.TrimRight(results[0][5], "*"))
				if t != nil {
					g := repository.GetGroupByID(*t.GroupId)
					if g != nil && *g.CategoryId == 8 {
						item.ChargeTypeId = *t.Id

						items = append(items, entity.Item{
							TypeName: strings.TrimRight(results[0][5], "*"),
							TypeId: *t.Id,
							GroupId: *t.GroupId,
							Quantity: 1,
						})
					}
				}
			}

			items = append(items, item)
		}
	}

	return items, nil
}
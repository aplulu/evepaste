package parser

import (
	"regexp"
	"fareastdominions.com/evepaste/eve/entity"
	"encoding/json"
	"fareastdominions.com/evepaste/eve/fit"
	"strings"
)


var CLFRegex *regexp.Regexp = regexp.MustCompile(`^{.*clf-version`)

type CLFParser struct {
}

func (p *CLFParser) Type() string {
	return "fit"
}

func (p *CLFParser) Parse(lines []string) ([]entity.Item, error) {
	items := []entity.Item{}

	text := strings.TrimSpace(strings.Join(lines, ""))

	if !CLFRegex.MatchString(text) {
		return items, nil
	}

	clf := fit.CLF{}
	err := json.Unmarshal([]byte(text), &clf)
	if err != nil {
		return items, err
	}

	name := "No name"
	if clf.Metadata.Title != "" {
		name = clf.Metadata.Title
	}

	if clf.Ship.TypeId > 0 {
		items = append(items, entity.Item{
			TypeId: clf.Ship.TypeId,
			Quantity: 1,
			Name: name,
			SlotType: "ship",
		})
	}

	if len(clf.Presets) > 0 {
		preset := clf.Presets[0]
		for _, m := range preset.Modules {
			item := entity.Item{
				TypeId: m.TypeId,
				Quantity: 1,
			}


			// Charges
			if len(m.Charges) > 0 {
				item.ChargeTypeId = m.Charges[0].TypeId

				items = append(items, entity.Item{
					TypeId: m.Charges[0].TypeId,
					Quantity: m.Charges[0].Quantity,
				})
			}

			items = append(items, item)
		}
	}

	if len(clf.Drones) > 0 {
		for _, m := range clf.Drones[0].InBay {
			items = append(items, entity.Item{
				TypeId: m.TypeId,
				Quantity: m.Quantity,
				SlotType: "drone",
			})
		}

		for _, m := range clf.Drones[0].InSpace {
			items = append(items, entity.Item{
				TypeId: m.TypeId,
				Quantity: m.Quantity,
				SlotType: "drone",
			})
		}
	}

	return items, nil
}
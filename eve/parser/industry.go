package parser

import (
	"regexp"
	"strconv"
	"strings"
	"fareastdominions.com/evepaste/eve/entity"
)


/*
Blueprint
Domination Control Tower Medium Blueprint
Reprocessed materials
Tritanium (1750000 Units)
Pyerite (291667 Units)
Mexallon (87500 Units)
Isogen (29167 Units)
Nocxium (14583 Units)
Zydrine (23334 Units)
Megacyte (8748 Units)
Broadcast Node (10 Units)
Integrity Response Drones (18 Units)
Nano-Factory (16 Units)
Organic Mortar Applicators (16 Units)
Recursive Computing Module (10 Units)
Self-Harmonizing Power Core (10 Units)
Sterile Conduits (16 Units)
Wetware Mainframe (7 Units)
Capital Construction Parts (4 Units)


ストラクチャ オフィスセンター*（10 ユニット）
ストラクチャ 広報ネクサス*（4 ユニット）
*/

var IndustryRegex *regexp.Regexp = regexp.MustCompile(`^([\S ]+?)(\s\((\d+) Units?\)|（(\d+) ユニット）)?$`)

type IndustryParser struct {
}

func (p *IndustryParser) Type() string {
	return "industry"
}

func (p *IndustryParser) Parse(lines []string) ([]entity.Item, error) {
	items := []entity.Item{}

	for _, line := range lines {
		results := IndustryRegex.FindAllStringSubmatch(strings.TrimSpace(line), -1)

		if len(results) == 1 {
			var q int = 1
			if results[0][3] != "" {
				quantity, err := strconv.Atoi(results[0][3])
				if err != nil {
					continue
				}

				q = quantity
			} else if results[0][4] != "" {
				quantity, err := strconv.Atoi(results[0][4])
				if err != nil {
					continue
				}

				q = quantity
			} else {
				q = 1
			}

			items = append(items, entity.Item{
				TypeName: strings.TrimRight(results[0][1], "*"),
				Quantity: q,
			})
		}
	}

	return items, nil
}
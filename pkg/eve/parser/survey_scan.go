package parser

import (
	"regexp"
	"strings"
	"github.com/evepaste/evepaste/pkg/eve/entity"
)


var surveyScanRegex *regexp.Regexp = regexp.MustCompile(`^([\S ]+)\t([\d,]+)\t(([\d,\.]+)\s?(AU|km|m)|\-)$`)

type SurveyScanParser struct {
	BaseParser
}

func (p *SurveyScanParser) Type() string {
	return "survey_scan"
}

func (p *SurveyScanParser) Parse(lines []string) ([]entity.Item, error) {
	items := []entity.Item{}

	for _, line := range lines {
		results := surveyScanRegex.FindAllStringSubmatch(strings.TrimSpace(line), -1)
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
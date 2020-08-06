package parser

import (
	"regexp"
	"strings"
	"github.com/evepaste/evepaste/pkg/eve/entity"
)


var ContractRegex *regexp.Regexp = regexp.MustCompile(`^([\S ]+)\t([\d,\.]+)\t([\S ]+)\t([\S ]+)(\t([\S ]+))?$`)

type ContractParser struct {
	BaseParser
}

func (p *ContractParser) Type() string {
	return "contract"
}

func (p *ContractParser) Parse(lines []string) ([]entity.Item, error) {
	items := []entity.Item{}

	for _, line := range lines {
		results := ContractRegex.FindAllStringSubmatch(strings.TrimSpace(line), -1)

		if len(results) == 1 {
			quantity, err := p.parseInt(results[0][2])
			if err != nil {
				continue
			}


			items = append(items, entity.Item{
				TypeName: strings.TrimRight(results[0][1], "*"),
				GroupName: strings.TrimRight(results[0][3], "*"),
				Quantity: quantity,
			})
		}
	}

	return items, nil
}
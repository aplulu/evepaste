package parser

import (
	"regexp"
	"strings"
	"fareastdominions.com/evepaste/eve/entity"
	"strconv"
)


var DscanRegex *regexp.Regexp = regexp.MustCompile(`^([\S ]+)\t([\S ]+)\t(([\d,\.]+)\s?(AU|km|m)|\-)$`)

type DscanParser struct {
}

func (p *DscanParser) Type() string {
	return "dscan"
}

func (p *DscanParser) Parse(lines []string) ([]entity.Item, error) {
	items := []entity.Item{}

	for _, line := range lines {
		results := DscanRegex.FindAllStringSubmatch(strings.TrimSpace(line), -1)
		if len(results) == 1 {
			var distance int64
			if results[0][5] == "AU" {
				f, err := strconv.ParseFloat(results[0][4], 64)
				if err != nil {
					continue
				}
				distance = int64(f * 149597870700)
			} else if results[0][5] == "km" {
				d, err := strconv.Atoi(strings.Replace(results[0][4], ",", "", -1))
				if err != nil {
					continue
				}
				distance = int64(d * 1000)
			} else if results[0][5] == "m" {
				d, err := strconv.Atoi(strings.Replace(results[0][4], ",", "", -1))
				if err != nil {
					continue
				}
				distance = int64(d)
			} else {
				distance = -1
			}

			items = append(items, entity.Item{
				TypeName: strings.TrimRight(results[0][2], "*"),
				Quantity: 1,
				Name: results[0][1],
				Distance: distance,
			})
		}
	}

	return items, nil
}
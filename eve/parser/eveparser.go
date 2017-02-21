package parser

import (
	"strings"
	"regexp"
	"strconv"
	"fareastdominions.com/evepaste/eve/entity"
	"fareastdominions.com/evepaste/eve/repository"
)

type Parser interface {
	Type() string
	Parse(lines []string) ([]entity.Item, error)
}

type BaseParser struct {

}

func (p *BaseParser) parseInt(v string) (int, error) {
	v = strings.Replace(v, ",", "", -1)
	parts := strings.SplitN(v, ".", 2)
	return strconv.Atoi(parts[0])
}

func Parse(text string) (string, []entity.Item) {
	lines := strings.Split(text, "\n")
	for i, _ := range lines {
		lines[i] = strings.TrimRight(lines[i], "\r\n")
	}

	parsers := []Parser{
		&ContractParser{},
		&InventoryParser{},
		&ViewContentsParser{},
		&ListParser{},
		&CargoScanParser{},
		&SurveyScanParser{},
		&DscanParser{},
		&PIParser{},
		&LootHistoryParser{},
		&EFTParser{},
		&DNAParser{},
		&CLFParser{},
		&IndustryParser{}, // アイテム名だけでマッチするので一番最後に
	}

	for _, v := range parsers {
		items, err := v.Parse(lines)
		if err != nil || len(items) == 0 {
			continue
		}

		items = resolveType(items, v.Type())

		return v.Type(), items
	}

	return "", nil
}

func resolveType(items []entity.Item, pasteType string) []entity.Item {
	results := make([]entity.Item, 0)
	for _, item := range items {
		var t *entity.InvType
		if item.TypeId > 0 {
			t = repository.GetTypeByID(item.TypeId)
		} else {
			t = repository.GetTypeByName(item.TypeName)
		}
		if t != nil {
			g := repository.GetGroupByID(*t.GroupId)
			if g != nil {
				item.GroupName = g.GetName("en")
				item.GroupId = *g.Id
			}

			//item.TypeName = *t.Name
			item.TypeId = *t.Id
			item.Volume = *t.Volume
			item.GroupId = *t.GroupId

			if pasteType == "fit" {
				if t.EffectId != nil {
					if *t.EffectId == 11 {
						item.SlotType = "low"
					} else if *t.EffectId == 12 {
						item.SlotType = "high"
					} else if *t.EffectId == 13 {
						item.SlotType = "med"
					} else if *t.EffectId == 2663 {
						item.SlotType = "rig"
					} else if *t.EffectId == 3772 {
						item.SlotType = "subsystem"
					}
				} else if *g.CategoryId == 18 {
					item.SlotType = "drone"
				}
			}

			results = append(results, item)
		}
	}
	return results
}

func ApplyI18nName(items []entity.Item, lang string) []entity.Item {
	lang = strings.SplitN(lang, "-", 2)[0]

	for i, item := range items {
		t := repository.GetTypeByID(item.TypeId)
		if t != nil {
			items[i].TypeName = t.GetName(lang)
			items[i].Volume = *t.Volume
		}

		if item.OriginalTypeId != 0 {
			t := repository.GetTypeByID(item.OriginalTypeId)
			if t != nil {
				items[i].OriginalTypeName = t.GetName(lang)
			}
		}
	}
	return items
}

var SunRegex = regexp.MustCompile(`^([\S ]+) - `)
var PlanetRegex = regexp.MustCompile(`^([\S ]+) [IVX]+`)
var CustomsRegex = regexp.MustCompile(`^Customs Office \(([\S ]+) [IVX]+\)$`)

func ScanSolarSystem(items []entity.Item) int32 {
	for _, item := range items {
		t := repository.GetTypeByID(item.TypeId)
		if t != nil {
			continue
		}

		if *t.GroupId == 6 { // Sun
			results := SunRegex.FindAllStringSubmatch(item.Name, -1)
			if len(results) == 1 && results[0][1] != "" {
				s := repository.GetSolarSystemByName(results[0][1])
				if s != nil {
					return *s.Id
				}
			}
		} else if *t.GroupId == 7 || *t.GroupId == 8 || *t.GroupId == 15 { // Planet, Moon, Station
			results := PlanetRegex.FindAllStringSubmatch(item.Name, -1)
			if len(results) == 1 && results[0][1] != "" {
				s := repository.GetSolarSystemByName(results[0][1])
				if s != nil {
					return *s.Id
				}
			}
		} else if *t.Id == 2233 { // Customs Office
			results := CustomsRegex.FindAllStringSubmatch(item.Name, -1)
			if len(results) == 1 && results[0][1] != "" {
				s := repository.GetSolarSystemByName(results[0][1])
				if s != nil {
					return *s.Id
				}
			}
		}

	}

	return 0
}

func ScanItems(items []entity.Item, systemId int, lang string) (result entity.ScanResult) {
	result = entity.ScanResult{}
	types := make(map[int32]*entity.ScanType, 0)
	groups := make(map[int32]*entity.ScanGroup, 0)

	for _, item := range items {
		t := repository.GetTypeByID(item.TypeId)
		if *t.Id == 0 {
			continue
		}
		g := repository.GetGroupByID(*t.GroupId)
		if *g.Id == 0 {
			continue
		}

		// ScanTypes
		if _, ok := types[*t.Id]; !ok {
			types[*t.Id] = &entity.ScanType{
				TypeId: *t.Id,
				TypeName: t.GetName(lang),
				GroupId: *g.Id,
				CategoryId: *g.CategoryId,
			}
		}

		// ScanGroups
		if _, ok := groups[*g.Id]; !ok {
			groups[*g.Id] = &entity.ScanGroup{
				GroupId: *g.Id,
				GroupName: g.GetName(lang),
				CategoryId: *g.CategoryId,
			}
		}

		if item.Distance != -1 && item.Distance <= 8000000 {
			types[*t.Id].OnGridCount++
			groups[*g.Id].OnGridCount++
			result.OnGridCount++
		} else {
			types[*t.Id].OffGridCount++
			groups[*g.Id].OffGridCount++
			result.OffGridCount++
		}
		types[*t.Id].Total++
		groups[*g.Id].Total++
		result.Total++
	}

	result.Types = make([]entity.ScanType, 0, len(types))
	for _, v := range types {
		result.Types = append(result.Types, *v)
	}

	result.Groups = make([]entity.ScanGroup, 0, len(groups))
	for _, v := range groups {
		result.Groups = append(result.Groups, *v)
	}

	if systemId > 0 {
		solarSystem := repository.GetSolarSystemByID(int32(systemId))
		if solarSystem != nil {
			result.SolarSystemName = solarSystem.GetName(lang)
			result.SolarSystemSecurity = *solarSystem.Security
			constellation := repository.GetConstellationByID(*solarSystem.ConstellationId)
			if constellation != nil {
				result.ConstellationName = constellation.GetName(lang)
			}
			region := repository.GetRegionByID(*solarSystem.RegionId)
			if region != nil {
				result.RegionName = region.GetName(lang)
			}
		}
	}

	return
}

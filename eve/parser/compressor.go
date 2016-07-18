package parser

import (
	"fareastdominions.com/evepaste/eve/entity"
	"fareastdominions.com/evepaste/eve/repository"
)

var ORE_COMPRESS_TABLE = map[int32]int32 {
	1230: 28432, // Veldsper
	17470: 28430, // Concentrated Veldspar
	17471: 28431, // Dense Veldspar

	1228: 28429, // Scordite
	17463: 28427, // Condensed Scordite
	17464: 28428, // Massive Scordite

	18: 28422, // Plagioclase
	17455: 28421,  // Azure Plagioclase
	17456: 28423, // Rich Plagioclase

	1224: 28424, // Pyroxeres
	17459: 28425, // Solid Pyroxeres
	17460: 28426, // Viscous Pyroxeres

	1227: 28416, // Omber
	17867: 28417, // Silvery Omber
	17868: 28415, // Golden Omber

	1226: 28406, // Jaspet
	17448: 28408, // Pure Jaspet
	17449: 28407, // Pristine Jaspet

	1231: 28403, // Hemorphite
	17444: 28405, // Vivid Hemorphite
	17445: 28404, // Radiant Hemorphite

	21: 28401, // Hedbergite
	17440: 28402, // Vitric Hedbergite
	17441: 28400, // Glazed Hedbergite

	1229: 28397, // Gneiss
	17865: 28398, // Iridescent Gneiss
	17866: 28399, // Prismatic Gneiss

	1232: 28394, // Dark Ochre
	17436: 28396, // Onyx Ochre
	17437: 28395, // Obsidian Ochre

	1225: 28391, // Crokite
	17432: 28393, // Sharp Crokite
	17433: 28392, // Crystalline Crokite

	19: 28420, // Spodumain
	17466: 28418, // Bright Spodumain
	17467: 28419, // Gleaming Spodumain

	1223: 28388, // Bistot
	17428: 28390, // Triclinic Bistot
	17429: 28389, // Monoclinic Bistot

	22: 28367, // Arkonor
	17425: 28385, // Crimson Arkonor
	17426: 28387, // Prime Arkonor

	11396: 28413, // Mercoxit
	17869: 28412, // Magma Mercoxit
	17870: 28414, // Vitreous Mercoxit

	20: 28410, // Kernite
	17452: 28411, // Luminous Kernite
	17453: 28471, // Fiery Kernite
};

var ICE_COMPRESS_TABLE = map[int32]int32 {
	16265: 28444, // White Glaze
	17976: 28441, // Pristine White Glaze

	16263: 28438, // Glacial Mass
	17977: 28442, // Smooth Glacial Mass

	16264: 28433, // Blue Ice
	17975: 28443, // Thick Blue Ice

	16262: 28434, // Clear Icicle
	17978: 28436, // Enriched Clear Icicle

	16266: 28439, // Glare Crust

	16267: 28435, // Dark Glitter

	16268: 28437, // Gelidus

	16269: 28440, // Krystallos
};

func CompressItems(items []entity.Item) []entity.Item {
	results := make([]entity.Item, 0)

	for _, item := range items {
		// ORE
		if compressedId, ok := ORE_COMPRESS_TABLE[item.TypeId]; ok {
			quantity := item.Quantity
			// あまり
			item.Quantity = quantity % 100
			if item.Quantity > 0 {
				results = append(results, item)
			}


			t := repository.GetTypeByID(compressedId)
			if t != nil {
				item.TypeName = t.GetName("en")
				item.TypeId = *t.Id
				item.Volume = *t.Volume
				item.GroupId = *t.GroupId
				item.Quantity = int(float64(quantity) * 0.01)
			}

			g := repository.GetGroupByID(item.GroupId)
			if g != nil {
				item.GroupName = g.GetName("en")
				item.GroupId = *g.Id
			}
		// ICE
		} else if compressedId, ok := ICE_COMPRESS_TABLE[item.TypeId]; ok {
			t := repository.GetTypeByID(compressedId)
			if t != nil {
				item.TypeName = t.GetName("en")
				item.TypeId = *t.Id
				item.Volume = *t.Volume
				item.GroupId = *t.GroupId
			}

			g := repository.GetGroupByID(item.GroupId)
			if g != nil {
				item.GroupName = g.GetName("en")
				item.GroupId = *g.Id
			}
		}

		results = append(results, item)
	}

	return results
}

func DecompressItems(items []entity.Item) []entity.Item {
	results := make([]entity.Item, 0)

	for _, item := range items {
		// ORE
		if decompressedId := getDecompressedOreTypeId(item.TypeId); decompressedId != 0 {
			t := repository.GetTypeByID(decompressedId)
			if t != nil {
				//item.TypeName = *t.Name
				item.TypeId = *t.Id
				item.Volume = *t.Volume
				item.GroupId = *t.GroupId
				item.Quantity = item.Quantity * 100
			}

			g := repository.GetGroupByID(item.GroupId)
			if g != nil {
				item.GroupName = g.GetName("en")
				item.GroupId = *g.Id
			}
			// ICE
		} else if decompressedId := getDecompressedIceTypeId(item.TypeId); decompressedId != 0 {
			t := repository.GetTypeByID(decompressedId)
			if t != nil {
				item.TypeName = t.GetName("en")
				item.TypeId = *t.Id
				item.Volume = *t.Volume
				item.GroupId = *t.GroupId
			}

			g := repository.GetGroupByID(item.GroupId)
			if g != nil {
				item.GroupName = g.GetName("en")
				item.GroupId = *g.Id
			}
		}

		results = append(results, item)
	}

	return results
}

func getDecompressedOreTypeId(typeId int32) int32 {
	for k, v := range ORE_COMPRESS_TABLE {
		if typeId == v {
			return k
		}
	}
	return 0
}

func getDecompressedIceTypeId(typeId int32) int32 {
	for k, v := range ICE_COMPRESS_TABLE {
		if typeId == v {
			return k
		}
	}
	return 0
}
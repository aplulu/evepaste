package repository

import (
	"time"
	"log"
	"encoding/gob"
	"os"
	"github.com/evepaste/evepaste/pkg/eve/entity"
	"io/ioutil"
)

var ID_TO_TYPE map[int32]*entity.InvType
var NAME_TO_TYPE_ID map[string]int32
var ID_TO_GROUP map[int32]*entity.InvGroup
var ID_TO_CATEGORY map[int32]*entity.InvCategory
var ID_TO_SOLARSYSTEM map[int32]*entity.SolarSystem
var NAME_TO_SOLARSYSTEM_ID map[string]int32
var ID_TO_CONSTELLATION map[int32]*entity.Constellation
var ID_TO_REGION map[int32]*entity.Region


func LoadTables(tableDir string) {
	start := time.Now()
	err := loadTablePB(tableDir)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("EVE tables loaded: %dms", int(time.Now().Sub(start).Nanoseconds() / int64(time.Millisecond)))
}

func loadTablePB(tableDir string) error {
	bytes, err := ioutil.ReadFile(tableDir + "data_set.pb")
	if err != nil {
		return err
	}

	dataSet := &entity.DataSet{}
	//err = proto.Unmarshal(bytes, dataSet)
	err = dataSet.Unmarshal(bytes)
	if err != nil {
		return err
	}

	// types
	ID_TO_TYPE = make(map[int32]*entity.InvType)
	NAME_TO_TYPE_ID = make(map[string]int32)
	for _, v := range dataSet.Types {
		ID_TO_TYPE[*v.Id] = v
		for _, t := range v.Names {
			NAME_TO_TYPE_ID[*t.Text] = *v.Id
		}
	}

	// groups
	ID_TO_GROUP = make(map[int32]*entity.InvGroup)
	for _, v := range dataSet.Groups {
		ID_TO_GROUP[*v.Id] = v
	}

	// category
	ID_TO_CATEGORY = make(map[int32]*entity.InvCategory)
	for _, v := range dataSet.Categories {
		ID_TO_CATEGORY[*v.Id] = v
	}

	// solar systems
	ID_TO_SOLARSYSTEM = make(map[int32]*entity.SolarSystem)
	NAME_TO_SOLARSYSTEM_ID = make(map[string]int32)
	for _, v := range dataSet.SolarSystems {
		ID_TO_SOLARSYSTEM[*v.Id] = v
		for _, t := range v.Names {
			NAME_TO_SOLARSYSTEM_ID[*t.Text] = *v.Id
		}
	}

	// constellations
	ID_TO_CONSTELLATION = make(map[int32]*entity.Constellation)
	for _, v := range dataSet.Constellations {
		ID_TO_CONSTELLATION[*v.Id] = v
	}

	// regions
	ID_TO_REGION = make(map[int32]*entity.Region)
	for _, v := range dataSet.Regions {
		ID_TO_REGION[*v.Id] = v
	}

	return nil
}

func loadTableGob(tableDir string) error {
	// types
	start := time.Now()
	file, err := os.Open(tableDir + "types.data")
	if err != nil {
		return err
	}
	log.Printf("types loaded: %dms", int(time.Now().Sub(start).Nanoseconds() / int64(time.Millisecond)))

	start = time.Now()
	types := make([]entity.InvType, 0)
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&types)
	file.Close()
	log.Printf("types decoded: %dms", int(time.Now().Sub(start).Nanoseconds() / int64(time.Millisecond)))

	start = time.Now()
	ID_TO_TYPE = make(map[int32]*entity.InvType)
	NAME_TO_TYPE_ID = make(map[string]int32)
	for _, v := range types {
		ID_TO_TYPE[*v.Id] = &v
		for _, t := range v.Names {
			NAME_TO_TYPE_ID[*t.Text] = *v.Id
		}
	}

	// groups
	file, err = os.Open(tableDir + "groups.data")
	if err != nil {
		return err
	}

	groups := make([]*entity.InvGroup, 0)
	decoder = gob.NewDecoder(file)
	err = decoder.Decode(&groups)
	file.Close()

	ID_TO_GROUP = make(map[int32]*entity.InvGroup)
	for _, v := range groups {
		ID_TO_GROUP[*v.Id] = v
	}

	// categories
	file, err = os.Open(tableDir + "categories.data")
	if err != nil {
		return err
	}

	categories := make([]*entity.InvCategory, 0)
	decoder = gob.NewDecoder(file)
	err = decoder.Decode(&categories)
	file.Close()

	ID_TO_CATEGORY = make(map[int32]*entity.InvCategory)
	for _, v := range categories {
		ID_TO_CATEGORY[*v.Id] = v
	}

	// solarSystems
	file, err = os.Open(tableDir + "solarsystems.data")
	if err != nil {
		return err
	}

	solarSystems := make([]*entity.SolarSystem, 0)
	decoder = gob.NewDecoder(file)
	err = decoder.Decode(&solarSystems)
	file.Close()

	ID_TO_SOLARSYSTEM = make(map[int32]*entity.SolarSystem)
	NAME_TO_SOLARSYSTEM_ID = make(map[string]int32)
	for _, v := range solarSystems {
		ID_TO_SOLARSYSTEM[*v.Id] = v
		for _, t := range v.Names {
			NAME_TO_SOLARSYSTEM_ID[*t.Text] = *v.Id
		}
	}

	// constellations
	file, err = os.Open(tableDir + "constellations.data")
	if err != nil {
		return err
	}

	contellations := make([]*entity.Constellation, 0)
	decoder = gob.NewDecoder(file)
	err = decoder.Decode(&contellations)
	file.Close()

	ID_TO_CONSTELLATION = make(map[int32]*entity.Constellation)
	for _, v := range contellations {
		ID_TO_CONSTELLATION[*v.Id] = v
	}

	// regions
	file, err = os.Open(tableDir + "regions.data")
	if err != nil {
		return err
	}

	regions := make([]*entity.Region, 0)
	decoder = gob.NewDecoder(file)
	err = decoder.Decode(&regions)
	file.Close()

	ID_TO_REGION = make(map[int32]*entity.Region)
	for _, v := range regions {
		ID_TO_REGION[*v.Id] = v
	}

	return nil
}

func GetTypeByName(name string) *entity.InvType {
	id, ok := NAME_TO_TYPE_ID[name]
	if ok {
		return ID_TO_TYPE[id]
	} else {
		return nil
	}
}

func GetTypeByID(id int32) *entity.InvType {
	return ID_TO_TYPE[id]
}

func GetGroupByID(id int32) *entity.InvGroup {
	return ID_TO_GROUP[id]
}

func GetSolarSystemByID(id int32) *entity.SolarSystem {
	return ID_TO_SOLARSYSTEM[id]
}

func GetSolarSystemByName(name string) *entity.SolarSystem {
	id, ok := NAME_TO_SOLARSYSTEM_ID[name]
	if ok {
		return ID_TO_SOLARSYSTEM[id]
	} else {
		return nil
	}
}

func GetConstellationByID(id int32) *entity.Constellation {
	return ID_TO_CONSTELLATION[id]
}

func GetRegionByID(id int32) *entity.Region {
	return ID_TO_REGION[id]
}

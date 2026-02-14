package food

import (
	"fmt"
	"strings"
)

type Silo int           // values +, str +, marshaling -
type FoodBuildingType int   // values +, str +, marshaling -
type GlobalModifier int // values -, str -, marshaling -
type TileModifier int   // values -, str -, marshaling -

const (
	SiloNormal Silo = iota
	SiloUpgraded
)

const (
	FishCabin FoodBuildingType = iota
	HuntersLodge
	Farm
	StandardLure
	Sheepfold
)

var stringsToSilos = map[string]Silo{
	"normal":   SiloNormal,
	"upgraded": SiloUpgraded,
}

var silosStrings = map[Silo]string{
	SiloNormal:   "normal",
	SiloUpgraded: "upgraded",
}

func GetSilo(name string) (Silo, error) {
	name = strings.ToLower(name)
	silo, ok := stringsToSilos[name]
	if !ok {
		return 0, fmt.Errorf("no silo named %s", name)
	}
	return silo, nil
}

func (s Silo) String() string {
	str, ok := silosStrings[s]
	if !ok {
		return ""
	}
	return str
}

var stringsToFoodBuildings = map[string]FoodBuildingType{
	"fish_cabin":    FishCabin,
	"hunters_lodge": HuntersLodge,
	"farm":          Farm,
	"standard_lure": StandardLure,
	"sheepfold":     Sheepfold,
}

var foodBuildingsStrings = map[FoodBuildingType]string{
	FishCabin:    "fish_cabin",
	HuntersLodge: "hunters_lodge",
	Farm:         "farm",
	StandardLure: "standard_lure",
	Sheepfold:    "sheepfold",
}

func GetFoodBuilding(name string) (FoodBuildingType, error) {
	name = strings.ToLower(name)
	foodBuilding, ok := stringsToFoodBuildings[name]
	if !ok {
		return 0, fmt.Errorf("no silo named %s", name)
	}
	return foodBuilding, nil
}

func (s FoodBuildingType) String() string {
	str, ok := foodBuildingsStrings[s]
	if !ok {
		return ""
	}
	return str
}

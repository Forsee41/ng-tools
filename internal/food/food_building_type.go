package food

import (
	"fmt"
	"strings"

	"github.com/Forsee41/ng-tools/internal/constants"
)

type FoodBuildingType int // values +, str +, marshaling -

func (fbt FoodBuildingType) GetBaseProduction() float64 {
	switch fbt {
	case FishCabin:
		return constants.FishBaseProd
	case Farm:
		return constants.FieldBaseProd
	case HuntersLodge:
		return constants.DeerBaseProd
	}
	return 0
}

const (
	FishCabin FoodBuildingType = iota
	HuntersLodge
	Farm
)

var stringsToFoodBuildings = map[string]FoodBuildingType{
	"fish_cabin":    FishCabin,
	"hunters_lodge": HuntersLodge,
	"farm":          Farm,
}

var foodBuildingsStrings = map[FoodBuildingType]string{
	FishCabin:    "fish_cabin",
	HuntersLodge: "hunters_lodge",
	Farm:         "farm",
}

func GetFoodBuilding(name string) (FoodBuildingType, error) {
	name = strings.ToLower(name)
	foodBuilding, ok := stringsToFoodBuildings[name]
	if !ok {
		return 0, fmt.Errorf("no silo named %s", name)
	}
	return foodBuilding, nil
}

func (fbt FoodBuildingType) String() string {
	str, ok := foodBuildingsStrings[fbt]
	if !ok {
		return ""
	}
	return str
}

package food

import (
	"fmt"
	"strings"
)

type FoodBuildingType int // values +, str +, marshaling -

func (fbt FoodBuildingType) GetBaseProduction() float64 {
	switch fbt {
	case FishCabin:
		return FishBaseProd
	case Farm:
		return FieldBaseProd
	case HuntersLodge:
		return DeerBaseProd
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

type FoodTypeMultiCalculator func(gm *GlobalModifiers, tm *TileModifiers) float64

func (fbt FoodBuildingType) GetMultiCalculator() FoodTypeMultiCalculator {
	switch fbt {
	case FishCabin:
		return FishMultiCalculator
	case HuntersLodge:
		return DeerMultiCalculator
	case Farm:
		return FarmMultiCalculator
	}
	return nil
}

func FishMultiCalculator(gm *GlobalModifiers, tm *TileModifiers) float64 {
	return 0.0
}
func DeerMultiCalculator(gm *GlobalModifiers, tm *TileModifiers) float64 {
	return 0.0
}
func FarmMultiCalculator(gm *GlobalModifiers, tm *TileModifiers) float64 {
	return 0.0
}

func (fb *FoodBuilding) CalcWorkersTotalDebuff() float64 {
	totalDebuff := (float64(fb.InjuredAmount+fb.UnhappyAmount) * 0.2)
	return totalDebuff
}

func (fb *FoodBuilding) CalcModifiersMultiplier(
	gm *GlobalModifiers,
	tm *TileModifiers,
) float64 {
	calc := fb.BuildingType.GetMultiCalculator()
	result := calc(gm, tm)
	return result
}

func (fb *FoodBuilding) CalcProduction(gm *GlobalModifiers, tm *TileModifiers) float64 {
	baseProd := fb.BuildingType.GetBaseProduction()
	prod := ((fb.CalcModifiersMultiplier(gm, tm) * float64(fb.WorkersAmount)) -
		fb.CalcWorkersTotalDebuff()) * baseProd
	return prod
}

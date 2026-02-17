package food

import (
	"github.com/Forsee41/ng-tools/internal/constants"
)

type FoodBuildingTileUpgrades struct {
	FishUpgrades int
	FarmUpgrades int
	DeerUpgrades int
}

func (fbu *FoodBuildingTileUpgrades) GetTargetFoodUpgrades(fbt FoodBuildingType) int {
	switch fbt {
	case FishCabin:
		return fbu.FishUpgrades
	case Farm:
		return fbu.FarmUpgrades
	case HuntersLodge:
		return fbu.DeerUpgrades
	}
	return 0
}

type FoodTypeMultiCalculator func(
	gm *GlobalModifiers,
	tm *TileModifiers,
	foodBuildingTypeUpgrades int,
) float64

func (fbt FoodBuildingType) GetFoodSpecificMultiCalc() FoodTypeMultiCalculator {
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

func FishMultiCalculator(
	gm *GlobalModifiers,
	tm *TileModifiers,
	fishBuildingTileUpgrades int,
) float64 {
	return 0.0
}
func DeerMultiCalculator(
	gm *GlobalModifiers,
	tm *TileModifiers,
	deerBuildingTileUpgrades int,
) float64 {
	return 0.0
}
func FarmMultiCalculator(gm *GlobalModifiers,
	tm *TileModifiers,
	farmBuildingTileUpgrades int,
) float64 {
	return 0.0
}

func CalcCommonModifiersMultiplier(gm *GlobalModifiers, tm *TileModifiers, upgraded bool) float64 {
	// Substracts the base 1.0
	result := 0.0

	if gm.SqrlWinterPrep {
		result += constants.SqrlWinterPrepMulti
	}

	if gm.FeastGoatRelic {
		result += constants.FeastGoatRelicMuli
	} else if gm.FeastGoat {
		result += constants.FeastGoatMuli
	} else if gm.Feast {
		result += constants.FeastMuli
	}

	if gm.Overwork {
		result += constants.OverworkMulti
	}

	result += float64(gm.Stag500Stacks) * constants.Stag500StackMulti

	if gm.ClanBonuses.Ox200Buff {
		result += float64(gm.OxMilitaryUnits) * constants.OxMilitaryUnitMulti
	}

	if gm.ClanBonuses.KrakenImplicit && gm.Kraken5Wyrd {
		result += constants.Kraken5WyrdMulti
	}

	if gm.JarActive && gm.Happiness > 0 {
		happinessCounted := 0.0
		if gm.Happiness > constants.JarMaxHappiness {
			happinessCounted = 10.0
		} else {
			happinessCounted = gm.Happiness
		}
		result += happinessCounted * constants.JarPerHappinessMulti
	}

	if gm.ClanBonuses.Bear200Buff && tm.Bear200Tile {
		result += constants.Bear200BuffMulti
	}

	if gm.ClanBonuses.Horse500Buff && tm.HorseRelicNearby {
		result += constants.Horse500BuffMulti
	}

	if gm.ClanBonuses.Kraken200Buff && tm.SpectralWarrior {
		result += constants.KrakenSpectralMulti
	}

	if gm.ClanBonuses.OxImplicit {
		result += constants.OxImplicitMulti
	}
	if tm.CityBulder {
		result += constants.CityBuilderMulti
	}

	result += float64(tm.Silos) * constants.SiloMulti
	result += float64(tm.SilosUpgraded) * constants.UpgradedSiloMulti
	if gm.Lores.Eradication {
		result += float64(tm.Silos+tm.SilosUpgraded) * constants.EradicationBonusMulti
	}

	if tm.FertileSoils {
		result += constants.FertileSoilsMulti
	}

	result += float64(tm.VerdandiStacks) * constants.VerdandiStackMulti

	if gm.Lores.Herbalism {
		result += float64(tm.MendersActive) * constants.MenderBuffMulti
	}

	if gm.Lores.LayOfTheLand && tm.LayOfTheLandTile {
		result += constants.LayOfTheLandMulti
	}

	if !upgraded && gm.Lores.Handiwork {
		result += constants.HandiworkMulti
	}

	if tm.DepletedField {
		result -= constants.DepletedFieldPenaltyMulti
	}

	if gm.Lores.AdvancedProspection {
		result += float64(tm.AdvancedProspectionStacks) * constants.AdvancedProspectionStackMulti
	}

	return result

}

func (fb *FoodBuilding) CalcWorkersTotalDebuff(dragonImplicit bool) float64 {
	totalDebuff := float64(fb.InjuredAmount) * constants.InjuredPenaltyMulti
	if !dragonImplicit {
		totalDebuff += float64(fb.UnhappyAmount) * constants.UnhappyPenaltyMulti
	}
	return totalDebuff
}

func (fb *FoodBuilding) CalcModifiersMultiplier(
	gm *GlobalModifiers,
	tm *TileModifiers,
	fbu *FoodBuildingTileUpgrades,
) float64 {
	foodSpecificCalc := fb.BuildingType.GetFoodSpecificMultiCalc()
	buildingUpgrades := fbu.GetTargetFoodUpgrades(fb.BuildingType)
	result := 1.0
	result += CalcCommonModifiersMultiplier(gm, tm, fb.Upgraded)
	result += foodSpecificCalc(gm, tm, buildingUpgrades)
	return result
}

func (fb *FoodBuilding) CalcProduction(
	gm *GlobalModifiers,
	tm *TileModifiers,
	fbu *FoodBuildingTileUpgrades,
) float64 {
	baseProd := fb.BuildingType.GetBaseProduction()
	prodMulti := fb.CalcModifiersMultiplier(gm, tm, fbu) * float64(fb.WorkersAmount)
	prodMulti -= fb.CalcWorkersTotalDebuff(gm.ClanBonuses.DragonImplicit)
	return prodMulti * baseProd
}

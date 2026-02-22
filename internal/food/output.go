package food

type CalcOutput struct {
	TotalProd     float64
	PerPop        float64
	TotalFish     FoodTypeProdOutput
	TotalDeer     FoodTypeProdOutput
	TotalFarm     FoodTypeProdOutput
	UpgradeValues UpgradeValues
}

type UpgradeValues struct {
	BuildingUpgrades             []BuildingUpgradeValue
	SilosValues                  []SiloValue
	SiloUpgradeValues            []SiloUpgradeValue
	ForgeValues                  ForgeValues
	EradicationSiloUpgradesValue float64 // total value of era + all silo upgrades

	// values of all other bonuses will go here
}

type ForgeValues struct {
	DeerForgeValue float64
	FishForgeValue float64
	FarmForgeValue float64
}

type BuildingUpgradeValue struct {
	Id    int
	Value float64
}

type SiloValue struct {
	TileId int
	SiloId int
	Value  float64
}

type SiloUpgradeValue struct {
	TileId         int
	SiloId         int
	UpgradeValue   float64
	TotalSiloValue float64
}

type FoodTypeProdOutput struct {
	TotalProd   float64
	Workers     float64
	PerPop      float64
	PercentProd float64
}

type TileCalcOutput struct {
	TotalProd           float64
	PerFoodBuildingProd []FoodBuildingProdOutput
}

type FoodBuildingProdOutput struct {
	Id         int
	Production float64
	PerPop     float64
}

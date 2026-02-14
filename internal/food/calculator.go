package food

type FoodCalculator struct {
	I FoodCalcInput // input
}

func NewFoodCalculator(input FoodCalcInput) FoodCalculator {
	return FoodCalculator{I: input}
}

func (fc *FoodCalculator) CalcWorkersAmount() int {
	result := 0
	for _, tile := range fc.I.Tiles {
		for _, building := range tile.FoodBuildings {
			result += building.WorkersAmount
		}
	}
	return result
}

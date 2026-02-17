package food

type FoodCalcInput struct {
	GlobalModifiers GlobalModifiers
	Tiles           []Tile
}

type GlobalModifiers struct {
	SqrlWinterPrep    bool // done
	Winter            bool
	Forges            Forges
	Feast             bool    // done
	FeastGoat         bool    // done
	FeastGoatRelic    bool    // done
	Overwork          bool    // done
	Stag500Stacks     int     // # of stacks of 200 fame // done
	OxMilitaryUnits   int     // done
	Kraken5Wyrd       bool    // done
	Happiness         float64 // done
	JarActive         bool    // done
	Lores             Lores
	ClanBonuses       ClanBonuses
	ValhallaEventBuff bool // idk
}

type Tile struct {
	Modifiers     TileModifiers
	FoodBuildings []FoodBuilding
}

type TileModifiers struct {
	Silos                     int  // done
	SilosUpgraded             int  // done
	CityBulder                bool // done
	FertileSoils              bool // done
	Bear200Tile               bool // done
	VerdandiStacks            int  // done
	MendersActive             int  // done
	DepletedField             bool // done, TODO: check if additive
	HorseRelicNearby          bool // done
	SpectralWarrior           bool // done
	SqrlEldritchBuff          bool // TODO: check mechanics
	AdvancedProspectionStacks int  // done
	LayOfTheLandTile          bool
}

type FoodBuilding struct {
	BuildingType   FoodBuildingType
	WorkersAmount  int
	SlavesAmount   int
	UnhappyAmount  int // done
	InjuredAmount  int // done
	BrawlersAmount int
	Upgraded       bool
}

type Forges struct {
	Farm bool
	Fish bool
	Deer bool
}

type Lores struct {
	Hearthstone         bool
	Freya               bool
	BoarFreya           bool
	Eradication         bool // done
	Harpoons            bool
	Herbalism           bool // done
	Handiwork           bool // done
	BloodSweatTears     bool // thralls give bonus prod
	VolundFire          bool // prod forges are more efficient
	AdvancedProspection bool // done
	LayOfTheLand        bool // done
}

type ClanBonuses struct {
	BearImplicit     bool // no winter penalties
	Bear200Buff      bool // done
	DragonImplicit   bool // done
	Horse500Buff     bool // done
	KrakenImplicit   bool // fish is more efficient
	Kraken200Buff    bool // done
	OxImplicit       bool // done
	Ox200Buff        bool // done
	LynxImplicit     bool // hunter prod buff
	SquirrelImplicit bool // done winter prep, TODO: winter
}

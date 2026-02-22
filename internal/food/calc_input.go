package food

type FoodCalcInput struct {
	GlobalModifiers GlobalModifiers
	Tiles           []Tile
}

type GlobalModifiers struct {
	SqrlWinterPrep    bool // done
	Winter            bool
	Forges            Forges  // done
	Feast             bool    // done
	FeastGoat         bool    // done
	FeastGoatRelic    bool    // done
	Overwork          bool    // done
	Stag500Stacks     int     // done
	OxMilitaryUnits   int     // done
	Kraken5Wyrd       bool    // done
	Happiness         float64 // done
	JarActive         bool    // done
	Lores             Lores
	ClanBonuses       ClanBonuses
	ValhallaEventBuff bool // idk
}

type Tile struct {
	Id            int
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
	LayOfTheLandTile          bool // done
}

type FoodBuilding struct {
	Id             int
	BuildingType   FoodBuildingType
	WorkersAmount  int
	SlavesAmount   int
	UnhappyAmount  int // done
	InjuredAmount  int // done
	BrawlersAmount int
	Upgraded       bool // done
}

type Forges struct {
	Farm bool // done
	Fish bool // done
	Deer bool // done
}

type Lores struct {
	Hearthstone         bool
	Freya               bool
	BoarFreya           bool
	Eradication         bool // done
	Harpoons            bool // done
	Herbalism           bool // done
	Handiwork           bool // done
	BloodSweatTears     bool // thralls give bonus prod
	VolundFire          bool // done
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
	LynxImplicit     bool // done
	SquirrelImplicit bool // done winter prep, TODO: winter
}

package food

type CalcInput struct {
	GlobalModifiers GlobalModifiers
	Tiles           []Tile
}

type GlobalModifiers struct {
	SqrlWinterPrep  bool
	Winter          bool
	Forges          Forges
	Feast           bool
	FeastGoat       bool
	FeastGoatRelic  bool
	Overwork        bool
	Stag500Stacks   int // # of stacks of 200 fame
	OxMilitaryUnits int
	Lores           Lores
	ClanBonuses     ClanBonuses
}

type Tile struct {
	Modifiers        TileModifiers
	FoodBuildings    []FoodBuilding
	LayOfTheLandTile bool
}

type TileModifiers struct {
	Silos                     int
	SilosUpgraded             int
	CityBulder                bool
	FertileSoils              bool
	VerdandiStacks            int
	MendersActive             int
	DepletedField             bool // snake after burn
	HorseRelicNearby          bool
	SpectralWarrior           bool
	SqrlEldritchBuff          bool // idk
	AdvancedProspectionStacks int
}

type FoodBuilding struct {
	BuildingType   FoodBuildingType
	WorkersAmount  int
	SlavesAmount   int
	UnhappyAmount  int
	StarvingAmount int
	FreezingAmount int
	BrawlersAmount int
	Upgraded       bool
}

type Forges struct {
	Farm  bool
	Fish  bool
	Deers bool
}

type Lores struct {
	Hearthstone         bool
	Freya               bool
	BoarFreya           bool
	Eradication         bool
	Harpoons            bool
	WinterFestival      bool
	Herbalism           bool
	Handiwork           bool
	BloodSweatTears     bool // thralls give bonus prod
	VolundFire          bool // prod forges are more efficient
	AdvancedProspection bool
	LayOfTheLand        bool
}

type ClanBonuses struct {
	BearImplicit     bool // no winter penalties
	Bear200Buff      bool // chief/bear/camp bonus
	DragonImplicit   bool // no penalties
	Horse500Buff     bool // bonus prod to tiles near relic
	KrakenImplicit   bool // fish is more efficient
	Kraken200Buff    bool
	OxImplicit       bool
	Ox200Buff        bool
	LynxImplicit     bool // hunter prod buff
	SquirrelImplicit bool
}

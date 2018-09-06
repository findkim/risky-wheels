package main

const (
	// Named locations
	DustyDivot    = "Dusty Divot"
	FatalFields   = "Fatal Fields"
	FlushFactory  = "Flush Factory"
	GreasyGrove   = "Greasy Grove"
	HauntedHills  = "Haunted Hills"
	JunkJunction  = "Junk Junction"
	LazyLinks     = "Lazy Links"
	LonelyLodge   = "Loney Lodge"
	LootLake      = "Loot Lake"
	LuckyLanding  = "Lucky Landing"
	ParadisePalms = "Paradise Palms"
	PleasantPark  = "Pleasant Park"
	RetailRow     = "Retail Row"
	RiskyReels    = "Risky Reels"
	SaltySprings  = "Salty Springs"
	ShiftyShafts  = "Shifty Shafts"
	SnobbyShores  = "Snobby Shores"
	TiltedTowers  = "Tilted Towers"
	TomatoTown    = "Tomato Town"
	WailingWoods  = "Wailing Woods"

	// Unnamed locations
	BigChair         = "Big Chair"
	Containers       = "Containers"
	DeathMountain    = "Death Mountain"
	Factories        = "Factories"
	GhettoContainers = "Ghetto Containers"
	GusTrack         = "Gus Race Track"
	HeroHouse        = "Hero House"
	LittleMexico     = "Little Mexico"
	Motel            = "Motel"
	SoccerStadium    = "Soccer Stadium"
	TrailerPark      = "Trailer Park"
	VillainVilla     = "Villain Villa"
	VikingMountain   = "Viking Mountain"
)

var (
	NamedLocations = []string{
		DustyDivot,
		FatalFields,
		FlushFactory,
		GreasyGrove,
		HauntedHills,
		JunkJunction,
		LazyLinks,
		LonelyLodge,
		LootLake,
		LuckyLanding,
		ParadisePalms,
		PleasantPark,
		RetailRow,
		RiskyReels,
		SaltySprings,
		ShiftyShafts,
		SnobbyShores,
		TiltedTowers,
		TomatoTown,
		WailingWoods,
	}
	UnnamedLocations = []string{
		BigChair,
		Containers,
		DeathMountain,
		Factories,
		GhettoContainers,
		GusTrack,
		HeroHouse,
		LittleMexico,
		Motel,
		SoccerStadium,
		TrailerPark,
		VillainVilla,
		VikingMountain,
	}
	AllLocations   = append(NamedLocations, UnnamedLocations...)
	LocationCount  = len(NamedLocations) + len(UnnamedLocations)
	ImagesDir      = "public/images/"
	LocationImages = map[string]string{
		// Named
		DustyDivot:    "",
		FatalFields:   "",
		FlushFactory:  "",
		GreasyGrove:   "",
		HauntedHills:  "",
		JunkJunction:  "",
		LazyLinks:     "",
		LonelyLodge:   "",
		LootLake:      "",
		LuckyLanding:  "",
		ParadisePalms: "",
		PleasantPark:  "",
		RetailRow:     "",
		RiskyReels:    "",
		SaltySprings:  "season-5-map.jpg",
		ShiftyShafts:  "",
		SnobbyShores:  "",
		TiltedTowers:  "",
		TomatoTown:    "",
		WailingWoods:  "",

		// Unnamed
		BigChair:         "",
		Containers:       "",
		DeathMountain:    "",
		Factories:        "",
		GhettoContainers: "",
		GusTrack:         "",
		HeroHouse:        "",
		LittleMexico:     "",
		Motel:            "",
		SoccerStadium:    "",
		TrailerPark:      "",
		VillainVilla:     "",
		VikingMountain:   "",
	}
)

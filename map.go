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
	ChestnutChair      = "Chestnut Chair"
	ContainerCrypt     = "Container Crypt"
	CrustyContainers   = "Crusty Containers"
	DeathMountain      = "Death Mountain"
	ForgottenFactories = "Forgotten Factories"
	GusTrack           = "Gus Race Track"
	HeroHall           = "Hero Hall"
	LittleMexico       = "Little Mexico"
	ModestMotel        = "Modest Motel"
	SoccerStadium      = "Soccer Stadium"
	TrailerPark        = "Trailer Park"
	VillainVilla       = "Villain Villa"
	VikingVillage      = "Viking Village"
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
		ChestnutChair,
		ContainerCrypt,
		CrustyContainers,
		DeathMountain,
		ForgottenFactories,
		GusTrack,
		HeroHall,
		LittleMexico,
		ModestMotel,
		SoccerStadium,
		TrailerPark,
		VillainVilla,
		VikingVillage,
	}
	AllLocations  = append(NamedLocations, UnnamedLocations...)
	LocationCount = len(NamedLocations) + len(UnnamedLocations)
	ImagesDir     = "public/images/"
	// LocationImages maps location to image file with map coordinates
	// currently only supports unnamed locations
	LocationImages = map[string]string{
		ChestnutChair:      "chestnut-chair.png",
		ContainerCrypt:     "container-crypt.png",
		DeathMountain:      "death-mountain.png",
		ForgottenFactories: "forgotten-factories.png",
		CrustyContainers:   "crusty-containers.png",
		GusTrack:           "gus-track.png",
		HeroHall:           "hero-hall.png",
		LittleMexico:       "little-mexico.png",
		ModestMotel:        "modest-motel.png",
		SoccerStadium:      "soccer-stadium.png",
		TrailerPark:        "trailer-park.png",
		VillainVilla:       "villain-villa.png",
		VikingVillage:      "viking-village.png",
	}
)

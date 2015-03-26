package types

type World struct {
	DayTime bool
	Time float32
	MoonPhase int
	BloodMoon bool
	PumpkinMoon bool
	SnowMoon bool
	Eclipse bool
	MoonType int

	HardMode bool
	ServerSideCharacter bool

	LeftWorld float32
	RightWorld float32
	TopWorld float32
	BottomWorld float32
	MaxTilesX int
	MaxTilesY int

	SpawnTileX int
	SpawnTileY int

	WorldSurface float64
	RockLayer float64

	WorldID int
	WorldName string

	TreeX [4]int
	TreeStyle [4]int
	CaveBackX [4]int
	CaveBackStyle [4]int

	IceBackStyle int
	JungleBackStyle int
	HellBackStyle int

	MaxRaining float32
	Raining bool
	WindSpeedSet float32

	CloudLimit int
	NumClouds int

	CloudBGActive float32
}
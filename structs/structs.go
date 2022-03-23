package structs

// Ant
type Ant struct {
	Id          int
	Path        []*Room
	CurrentRoom *Room
	RoomsPassed int
}

// Room in farm
type Room struct {
	Name    string
	Ants    int
	X_pos   int
	Y_pos   int
	IsStart bool
	IsEnd   bool
	Links   []*Room
}

// Data to generate future farm
type GenerationData struct {
	NumberOfAnts int
	Rooms        []string
	Links        []string
	StartIndex   int
	EndIndex     int
}

// Structurized paths by links in start room
type PathStuct struct {
	PathName string
	Paths    [][]*Room
}

var (
	ANTCOUNTER  string // Amount of ants to spawn
	STARTROOMID int
	ENDROOMID   int
	FARM        []Room // Farm
)

var (
	BEST_TURNS_RES        int
	BEST_PATH             [][]*Room
	BEST_ROOMS_IN_USE_RES int
)

package structs

// Farm represents an adjaceny list graph
type Farm struct {
	Ants  []*Ant
	Rooms []*Room
}

// Ant
type Ant struct {
	Id          int
	CurrentRoom *Room
	RoomsPassed int
}

// Room in farm (Vertex in the graph)
type Room struct {
	Name    string
	Ants    int
	IsStart bool
	IsEnd   bool
	Links []*Room
}

// Data to generate future farm
type GenerationData struct {
	NumberOfAnts int
	Rooms        []string
	Links        []string
	StartIndex   int
	EndIndex     int
}

// Structurised paths by links in start room
type PathStruct struct {
	Path []*Room
	Ants []*Ant
}
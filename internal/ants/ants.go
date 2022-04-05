package ants

import (
	"fmt"
)

type Room struct {
	Name    string
	Ants    int
	IsStart bool
	IsEnd   bool
	Links   []string
}
type Ant struct {
	Id          int
	Path        []string
	RoomsPassed int
}
type Path struct {
	Id          int
	RoomsInPath []string
}

func CreateAnts() Ant {
	rooms := []Room{
		{

			Name:    "1",
			Ants:    0,
			IsStart: true,
			IsEnd:   false,
			Links:   []string{"1", "0"},
		},
		{

			Name:    "0",
			Ants:    0,
			IsStart: false,
			IsEnd:   true,
			Links:   []string{"0", "1"},
		},
	}

	paths := []Path{
		{
			Id:          1,
			RoomsInPath: []string{"1", "0"},
		},
	}

	noOfPaths := lenPathStruct(paths)

	for i := 1; i <= 3; i++ {
		var antToAdd Ant
		//var noOfAntsCurrentRoom int
		noOfRooms := 0
		antToAdd.Id = i

		var j int

		if i > noOfPaths {
			j = 0
		} else {
			j = i - 1
		}
		for _, p := range paths[j].RoomsInPath {
			for r := range rooms {
				if p == rooms[r].Name && rooms[r].IsEnd == true {
					antToAdd.Path = append(antToAdd.Path, rooms[r].Name)
					noOfRooms = noOfRooms + 1
					continue
				} else if p == rooms[r].Name {
					antToAdd.Path = append(antToAdd.Path, rooms[r].Name)
					noOfRooms = noOfRooms + 1
				}
			}
		}
		antToAdd.RoomsPassed = noOfRooms
		fmt.Println(antToAdd.Id, antToAdd.Path, antToAdd.RoomsPassed)

	}

	fmt.Println()

	return Ant{}
}

func lenRoomStruct(rooms []Room) int {
	lenRoomCounter := 0

	for _, room := range rooms {
		if room.Name != "" {
			lenRoomCounter = lenRoomCounter + 1
		}
	}

	return lenRoomCounter

}

func lenPathStruct(paths []Path) int {
	lenRoomCounter := 0

	for _, path := range paths {
		if path.Id != 0 {
			lenRoomCounter = lenRoomCounter + 1
		}
	}

	return lenRoomCounter

}

// Make a list of ants with their own path, current room and id
/*func CreateAnts(paths [][]*structs.Room) []structs.Ant {
	var result []structs.Ant
	for i := 1; i < structs.ANTCOUNTER+1; i++ {
		var antToAdd structs.Ant

		antToAdd.Id = i
		antToAdd.CurrentRoom = &structs.FARM[structs.STARTROOMID]
		antToAdd.Path = paths[i-1]

		result = append(result, antToAdd)
	}

	return result
}

// Create ants to move from start to end and print each ant step
func CreateStep(ants []structs.Ant) {
	var ANTCOUNTER int
	var passed bool = true
	for i := 0; i < len(ants); i++ {
		if ants[i].CurrentRoom.IsEnd {
			continue
		}
		nextRoomId := ants[i].RoomsPassed

		if ants[i].Path[nextRoomId].Ants != 0 {
			if !ants[i].Path[nextRoomId].IsEnd {
				continue
			}
		}

		ants[i].CurrentRoom.Ants--
		ants[i].CurrentRoom = ants[i].Path[nextRoomId]
		ants[i].CurrentRoom.Ants++
		ants[i].RoomsPassed++
		passed = false

		fmt.Print("L", ants[i].Id, "-", ants[i].CurrentRoom.Name, " ")
	}
	if passed && structs.FARM[structs.ENDROOMID].Ants == ANTCOUNTER {
		return
	} else {
		fmt.Println("")
		CreateStep(ants)
	}
}*/

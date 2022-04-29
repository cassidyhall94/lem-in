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

func CreateAnts() *Ant {
	rooms := []Room{
		{

			Name:    "0",
			Ants:    20,
			IsStart: true,
			IsEnd:   false,
			Links:   []string{"0", "1", "3"},
		},
		{

			Name:    "1",
			Ants:    0,
			IsStart: false,
			IsEnd:   false,
			Links:   []string{"1", "0", "2"},
		},
		{

			Name:    "2",
			Ants:    0,
			IsStart: false,
			IsEnd:   false,
			Links:   []string{"2", "1", "3"},
		},
		{

			Name:    "3",
			Ants:    0,
			IsStart: false,
			IsEnd:   true,
			Links:   []string{"3", "0", "2"},
		},
	}

	paths := []Path{
		{
			Id:          1,
			RoomsInPath: []string{"3"},
		},
		{
			Id:          2,
			RoomsInPath: []string{"1", "2", "3"},
		},
	}

	noOfPaths := lenPathStruct(paths)
	var j int

	antsPath := antsPerPath(paths)

	fmt.Println(antsPath)

	for i := 1; i <= 20; {
		var antToAdd *Ant = new(Ant)
		noOfRooms := 0

		if i > noOfPaths {
			j = 0
			noOfPaths = noOfPaths + 2
		}
		var pathToAdd []string
		for _, p := range paths[j].RoomsInPath {
			for r := range rooms {
				if p == rooms[r].Name && rooms[r].IsEnd == true {
					pathToAdd = append(pathToAdd, rooms[r].Name)
					noOfRooms = noOfRooms + 1
				} else if p == rooms[r].Name {
					pathToAdd = append(pathToAdd, rooms[r].Name)
					noOfRooms = noOfRooms + 1
				}
			}
		}
		j = j + 1
		antToAdd.Id = i
		antToAdd.Path = pathToAdd
		antToAdd.RoomsPassed = noOfRooms

		fmt.Println(antToAdd)

		if i < 19 {
			i = i + 2
		} else {
			i = i + 1
		}
	}

	return &(Ant{})
}

func lenPathStruct(paths []Path) int {
	lenPathCounter := 0

	for _, path := range paths {
		if path.Id != 0 {
			lenPathCounter = lenPathCounter + 1
		}
	}

	return lenPathCounter

}

func diffRoomsInPaths(paths []Path) []int {
	difPaths := []int{}
	noOfRoomsPath := []int{}

	for n := 0; n < lenPathStruct(paths); n++ {
		noOfRooms := len(paths[n].RoomsInPath)
		noOfRoomsPath = append(noOfRoomsPath, noOfRooms)
	}

	for d := 1; d < len(noOfRoomsPath); d++ {
		difRooms := noOfRoomsPath[d] - noOfRoomsPath[d-1]
		difPaths = append(difPaths, difRooms)
	}

	return difPaths

}

func antsPerPath(paths []Path) []int {
	noOfPaths := lenPathStruct(paths)
	difPaths := diffRoomsInPaths(paths)
	antsPerPath := []int{}

	antsPerPath = append(antsPerPath, 20/noOfPaths+difPaths[0]-1)
	for dr := 1; dr < noOfPaths; dr++ {
		antCounter := 20 - antsPerPath[dr-1]
		if dr == len(difPaths) {
			antsPerPath = append(antsPerPath, antCounter/len(difPaths))
		} else {
			antsPerPath = append(antsPerPath, antCounter/len(difPaths)+difPaths[dr]-1)
		}
	}

	return antsPerPath
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
	}*/

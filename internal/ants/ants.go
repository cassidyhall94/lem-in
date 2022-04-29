package ants

import (
	"fmt"
	"sort"
)

type Room struct {
	Name    string
	Ants    int
	IsStart bool
	IsEnd   bool
	Links   []string
}
type Ant struct {
	Id   int
	Path []string
}
type Path struct {
	Id          int
	RoomsInPath []string
}

func CreateAnts() *Ant {

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

	antToAdd := make(map[int][]string)

	for p := 0; p < noOfPaths; p++ {

		var i int
		for i = p + 1; i <= 20-p; {

			antToAdd[i] = paths[p].RoomsInPath
			//antToAdd.Path = paths[p].RoomsInPath

			if i < 19 {
				i = i + 2
			} else {
				i = i + 1
			}
			//fmt.Println(antToAdd)
		}
	}

	//fmt.Println(antToAdd)

	keys := make([]int, 0, len(antToAdd))
	for k := range antToAdd {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	//fmt.Println(len(value))

	for i := 0; i < 3; i++ {
		for _, k := range keys {
			if len(antToAdd[k]) > i {
				fmt.Printf("L%v-%v ", k, antToAdd[k][i])
			} else {
				continue
			}
		}

		//fmt.Println(key, value[0])
	}

	return &Ant{}
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

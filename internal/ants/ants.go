package ants

import (
	"fmt"
	"sort"
	"strconv"
	//"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

type Room struct {
	Name    string
	Ants    int
	IsStart bool
	IsEnd   bool
	IsEmpty bool
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

func CreateAnts() []*Ant {
	rooms := []Room{
		{
			Name:    "1",
			Ants:    0,
			IsStart: false,
			IsEnd:   false,
			IsEmpty: true,
			Links:   []string{"1", "0", "2"},
		},
		{
			Name:    "2",
			Ants:    0,
			IsStart: false,
			IsEnd:   false,
			IsEmpty: true,
			Links:   []string{"2", "1", "3"},
		},
		{
			Name:    "3",
			Ants:    0,
			IsStart: false,
			IsEnd:   true,
			IsEmpty: true,
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

	var newAntToAdd = []*Ant{}

	for p := 0; p < noOfPaths; p++ {
		var i int
		for i = p + 1; i <= 20-p; {
			var antToAdd *Ant = new(Ant)

			antToAdd.Id = i
			antToAdd.Path = paths[p].RoomsInPath

			if i < 19 {
				i = i + 2
			} else {
				i = i + 1
			}

			newAntToAdd = append(newAntToAdd, antToAdd)
			sort.Slice(newAntToAdd, func(i, j int) bool {
				return newAntToAdd[i].Id < newAntToAdd[j].Id
			})
		}
	}

	for i := 0; i < len(newAntToAdd); i++ {
		for p := 0; p < len(newAntToAdd[i].Path); p++ {
			roomsName, _ := strconv.Atoi(newAntToAdd[i].Path[p])
			if roomsName-1 != 0 {
				rooms[roomsName-2].IsEmpty = *setTrue()
			}
			if rooms[roomsName-1].IsEmpty == true && rooms[roomsName-1].IsEnd == false {
				fmt.Printf("L%d-%s ", newAntToAdd[i].Id, newAntToAdd[i].Path[p])
				rooms[roomsName-1].IsEmpty = *setFalse()
				break
			} else if rooms[roomsName-1].IsEmpty == true && rooms[roomsName-1].IsEnd == true {
				fmt.Printf("L%d-%s ", newAntToAdd[i].Id, newAntToAdd[i].Path[p])
				break
			}
		}
	}

	return newAntToAdd
}

func setFalse() *bool {
	b := false
	return &b
}

func setTrue() *bool {
	b := true
	return &b
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

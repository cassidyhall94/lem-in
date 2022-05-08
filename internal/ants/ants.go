package ants

import (
	"fmt"
	"sort"
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

	ants := 0
	lastRoom := endRoom(rooms)
	//antToStartFrom := []int{}
	//antToStartFrom = append(antToStartFrom, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)

	if ants == 0 {
		for p := 0; p < noOfPaths; p++ {
			fmt.Printf("L%d-%s ", p+1, newAntToAdd[p].Path[0])
			if newAntToAdd[p].Path[0] != lastRoom {
				rooms[0].IsEmpty = *setFalse()
			}
		}
		ants = ants + noOfPaths
	}
	fmt.Println()
	if ants > 0 && ants <= 20 {
		emptyRooms := emptyRoomsInFarm(rooms)
		for i := 0; i < ants; i++ {
			if len(newAntToAdd[i].Path) < 2 {
				continue
			} else if newAntToAdd[i].Path[1] != "3" {
				emptyRooms = emptyRooms + 1
			}
		}
		fmt.Println(emptyRooms)

		for ant := 0; ant < emptyRooms; ant++ {
			room := 0
			roomOne := 0
			for r := 0; r < len(rooms); r++ {
				if rooms[r].Name == newAntToAdd[ant+1].Path[0] {
					room = r
				} else if len(newAntToAdd[ant+1].Path) > 1 && rooms[r].Name == newAntToAdd[ant+1].Path[1] {
					roomOne = r
				}
			}

			if rooms[room].IsEmpty == true && rooms[room].IsEnd == false {
				fmt.Printf("L%d-%s ", newAntToAdd[ant+1].Id, rooms[room].Name)
				rooms[room].IsEmpty = *setFalse()
			} else if len(newAntToAdd[ant+1].Path) > 1 && rooms[roomOne].IsEmpty == true && rooms[roomOne].IsEnd == false {
				rooms[room].IsEmpty = *setTrue()
				fmt.Printf("L%d-%s ", newAntToAdd[ant+1].Id, rooms[roomOne].Name)
				rooms[roomOne].IsEmpty = *setFalse()
			} else if rooms[room].IsEnd == true {
				fmt.Printf("L%d-%s ", newAntToAdd[ant+1].Id, rooms[room].Name)
				//antToStartFrom = RemoveIndex(antToStartFrom, ant)

			}
		}

		fmt.Println(emptyRooms)
		ants = ants + emptyRooms
		fmt.Println()
	}

	//fmt.Println(antToStartFrom)

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

func emptyRoomsInFarm(rooms []Room) int {
	emptyRooms := 0
	for r := 0; r < len(rooms); r++ {
		if rooms[r].IsEmpty == true {
			emptyRooms = emptyRooms + 1
		}
	}
	return emptyRooms
}

func endRoom(rooms []Room) string {
	endRoom := ""

	for r := 0; r < len(rooms); r++ {
		if rooms[r].IsEnd == true {
			endRoom = rooms[r].Name
		}
	}
	return endRoom
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

package ants

import (
	"fmt"
	//"strconv"
)

type Room struct {
	Name    string
	IsStart bool
	IsEnd   bool
	IsEmpty bool
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

func CreateAnts() ([]string, []int) {
	rooms := []Room{
		{
			Name:    "0",
			IsStart: true,
			IsEnd:   false,
			IsEmpty: true,
			Links:   []string{"0", "1", "3"},
		},
		{
			Name:    "1",
			IsStart: false,
			IsEnd:   false,
			IsEmpty: true,
			Links:   []string{"1", "0", "2"},
		},
		{
			Name:    "2",
			IsStart: false,
			IsEnd:   false,
			IsEmpty: true,
			Links:   []string{"2", "1", "3"},
		},
		{
			Name:    "3",
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

	antCounter := 5
	antsEndRoom := []int{}
	roomsTravelled := []string{}

	for i := 1; i <= antCounter; {
		for j := 0; j < lenPathStruct(paths); j++ {
			for _, p := range paths[j].RoomsInPath {
				for r := 0; r < len(rooms); r++ {
					rooms[r].IsEmpty = *setTrue()
					if p == rooms[r].Name && rooms[r].IsEnd == true {
						antsEndRoom = append(antsEndRoom, i)
						roomsTravelled = append(roomsTravelled, rooms[r].Name)
						i = i + 1
						break
					} else if p == rooms[r].Name && rooms[r].IsEmpty == true {
						rooms[r].IsEmpty = *setFalse()
						antsEndRoom = append(antsEndRoom, i)
						roomsTravelled = append(roomsTravelled, rooms[r].Name)
					}
				}
				continue
			}
		}
	}
	fmt.Println(antsEndRoom, roomsTravelled)

	return roomsTravelled, antsEndRoom
}

func setFalse() *bool {
	b := false
	return &b
}

func setTrue() *bool {
	b := true
	return &b
}

func remove(ants []int, item int) []int {
	newitems := []int{}

	for _, i := range ants {
		if i != item {
			newitems = append(newitems, i)
		}
	}

	return newitems
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

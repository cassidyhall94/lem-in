package ants

import (
	"lem-in/structs"
)

// Make a list of ants with their own path, current room and id
func CreateAnts(paths [][]*structs.Room) []structs.Ant {
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
	var passed bool = true
	for i := 0; i < len(ants); i++ {
		if ants[i].CurrentRoom.IsEnd {
			continue
		}
	}
}

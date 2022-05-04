package ants

import (
	"fmt"

	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

// func CreateAnts assigns a path to each ant
// (ants []*structs.Ant) is the same thing as (ants structs.farm) from the farm function
// ants will move through the shortest path, we can use current room and rooms passed to measure their movement
// loop through all the ants and move them a room, one ant at a time
// store information in variable to be printed at a later time
func MoveAnts(ants []*structs.Ant, path *structs.PathStruct) []*structs.Ant {
	for _, oneAntStruct := range ants {
		// fmt.Printf("oneAntStruct: %+v\n", oneAntStruct)
		for _, oneRoom := range path.Path {
			// fmt.Printf("oneRoom: %+v\n", oneRoom)
			if oneRoom.IsStart {
				oneAntStruct.CurrentRoom.Name = oneRoom.Name
				// fmt.Printf("oneAntStruct: %+v\n", oneAntStruct)
				// fmt.Printf("oneRoom: %+v\n", oneRoom)
			}
		}

	}

	return []*structs.Ant{}
}

// func lenPathStruct(paths []*structs.PathStruct) int {
// 	lenPathCounter := 0

// 	for _, path := range paths {
// 		for _, pathRoom := range path.Path {
// 			if pathRoom.Name != "" {
// 				lenPathCounter = lenPathCounter + 1
// 			}
// 		}
// 	}
// 	return lenPathCounter
// }

// func diffRoomsInPaths(paths []*structs.PathStruct) []int {
// 	difPaths := []int{}
// 	noOfRoomsPath := []int{}

// 	for n := 0; n < lenPathStruct(paths); n++ {
// 		noOfRooms := len(paths[n].Path)
// 		noOfRoomsPath = append(noOfRoomsPath, noOfRooms)
// 	}

// 	for d := 1; d < len(noOfRoomsPath); d++ {
// 		difRooms := noOfRoomsPath[d] - noOfRoomsPath[d-1]
// 		difPaths = append(difPaths, difRooms)
// 	}

// 	return difPaths
// }

// func antsPerPath(paths []*structs.PathStruct) []int {
// 	noOfPaths := lenPathStruct(paths)
// 	difPaths := diffRoomsInPaths(paths)
// 	antsPerPath := []int{}

// 	antsPerPath = append(antsPerPath, 20/noOfPaths+difPaths[0]-1)
// 	for dr := 1; dr < noOfPaths; dr++ {
// 		antCounter := 20 - antsPerPath[dr-1]
// 		if dr == len(difPaths) {
// 			antsPerPath = append(antsPerPath, antCounter/len(difPaths))
// 		} else {
// 			antsPerPath = append(antsPerPath, antCounter/len(difPaths)+difPaths[dr]-1)
// 		}
// 	}

// 	return antsPerPath
// }

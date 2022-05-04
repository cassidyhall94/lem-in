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

// Lx-y Lz-w Lr-o ...
//     x, z, r represents the ants numbers (going from 1 to number_of_ants) and y, w, o represents the rooms names.
//     A room is defined by "name coord_x coord_y", and will usually look like "Room 1 2", "nameoftheroom 1 6", "4 6 7".
//     The coordinates of the rooms are only important for the Bonus task: ant farm visualizer.
//     The links (tunnels) are defined by "name1-name2" and will usually look like "1-2", "2-5".
//     The L is for Lem-in.
// L1-2
// L1-3 L2-2
// L1-1 L2-3 L3-2
// L2-1 L3-3
// L3-1
func MoveAnts(ants []*structs.Ant, path *structs.PathStruct) []string {
	result := []string{}
	for _, oneAntStruct := range ants {
		// oneAntStruct: &{Id:1 CurrentRoom:<nil> RoomsPassed:0}
		for _, oneRoom := range path.Path {
			// oneRoom: &{Name:0 Ants:0 IsStart:true IsEnd:false Links:[]}
			if oneRoom.IsStart || oneRoom.IsEnd {
				oneRoom.Ants = len(ants)
				if !oneRoom.IsStart {
					oneAntStruct.RoomsPassed += 1
				}
				oneAntStruct.CurrentRoom = oneRoom
				// START: oneAntStruct.CurrentRoom: &{Name:0 Ants:3 IsStart:true IsEnd:false Links:[]}
				// END: oneAntStruct.CurrentRoom: &{Name:1 Ants:3 IsStart:false IsEnd:true Links:[]}

				// fmt.Printf("oneAntStruct.CurrentRoom: %+v\n", oneAntStruct.CurrentRoom)
			} else if !oneRoom.IsStart && !oneRoom.IsEnd {
				oneAntStruct.RoomsPassed += 1
				oneRoom.Ants = +1
				oneAntStruct.CurrentRoom = oneRoom
				// oneAntStruct.CurrentRoom: &{Name:3 Ants:1 IsStart:false IsEnd:false Links:[]}
				// oneAntStruct.CurrentRoom: &{Name:2 Ants:1 IsStart:false IsEnd:false Links:[]}

				// fmt.Printf("oneAntStruct.CurrentRoom: %+v\n", oneAntStruct.CurrentRoom)
				// fmt.Printf("oneAntStruct %+v\n", oneAntStruct)
			}
			fmt.Printf("oneAntStruct: %+v\n, CurrentRoom: %+v\n", oneAntStruct, oneAntStruct.CurrentRoom)
		}
	}
	return result
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

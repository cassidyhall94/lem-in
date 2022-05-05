package ants

import (
	"fmt"

	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

// func CreateAnts assigns a path to each ant

// Lx-y Lz-w Lr-o ...
//     x, z, r represents the ants numbers (going from 1 to number_of_ants) and y, w, o represents the rooms names.
//     A room is defined by "name coord_x coord_y", and will usually look like "Room 1 2", "nameoftheroom 1 6", "4 6 7".
//     The coordinates of the rooms are only important for the Bonus task: ant farm visualizer.
//     The links (tunnels) are defined by "name1-name2" and will usually look like "1-2", "2-5".
//     The L is for Lem-in.

// store everything to one string and then range over the string to print slices
// use roomsPassed to add conditional statements for formatting?
// fmt.Printf("oneAntStruct: %+v\n, CurrentRoom: %+v\n", oneAntStruct, oneAntStruct.CurrentRoom)
// fmt.Printf("Ant ID: %+v CurrentRoom: %+v RoomsPassed: %+v\n", oneAntStruct.Id, oneAntStruct.CurrentRoom.Name, oneAntStruct.RoomsPassed)
func MoveAnts(ants []*structs.Ant, path *structs.PathStruct) []string {
	result := []string{}
	for _, oneAntStruct := range ants {
		for _, oneRoom := range path.Path {
			if oneRoom.IsStart || oneRoom.IsEnd {
				oneRoom.Ants = len(ants)
				if oneRoom.IsEnd {
					oneAntStruct.RoomsPassed += 1
				}
				oneAntStruct.CurrentRoom = oneRoom
				if oneRoom.IsEnd {
					endPrint := fmt.Sprintf("L%v-%v", oneAntStruct.Id, oneAntStruct.CurrentRoom.Name)
					result = append(result, endPrint)
				}
			} else if !oneRoom.IsStart && !oneRoom.IsEnd {
				oneAntStruct.RoomsPassed += 1
				oneRoom.Ants = +1
				oneAntStruct.CurrentRoom = oneRoom
				roomPrint := fmt.Sprintf("L%v-%v", oneAntStruct.Id, oneAntStruct.CurrentRoom.Name)
				result = append(result, roomPrint)
			}
		}
		for _, resultString := range result {
			if oneAntStruct.RoomsPassed <= i {
			}
		}
		// fmt.Printf("result: %+v\n", result)

	}

	return result
}

// func sliceOfStrings(result []string) []string {
// 	oneAntStruct := &structs.Ant{}
// 	for i, resultString := range result {
// 		if oneAntStruct.RoomsPassed <= i {
// 			resultString = fmt.Sprintf("L%v-%v", oneAntStruct.Id, oneAntStruct.CurrentRoom.Name)
// 			result = append(result, resultString)
// 		}
// 	}
// 	return result
// }

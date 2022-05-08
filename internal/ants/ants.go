package ants

import (
	"fmt"
	"strings"

	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

// fmt.Printf("oneAntStruct: %+v\n, CurrentRoom: %+v\n", oneAntStruct, oneAntStruct.CurrentRoom)
// fmt.Printf("Ant ID: %+v CurrentRoom: %+v RoomsPassed: %+v\n", oneAntStruct.Id, oneAntStruct.CurrentRoom.Name, oneAntStruct.RoomsPassed)

// func AssignAnts assigns a path to each ant
// loop over paths to get each path, if the (length of rooms + number of ants currently in the path2) is greater than in path1, send the ant to path2, if not, send to path1
func AssignAnts(ants []*structs.Ant, paths []*structs.PathStruct) []*structs.PathStruct {
	for _, ant := range ants {
		for i, path := range paths {
			if (len(path.Path) + len(path.Ants)) > (len(paths[i+1].Path) + len(paths[i+1].Ants)) {

			}
		}
	}

	return []*structs.PathStruct{}
}

// func MoveAnts take the ants one by one and move them through their path
func MoveAnts(ants []*structs.Ant, path *structs.PathStruct) []string {
	result := make([]string, (len(ants) + len(path.Path) - 1))
	for i, oneAntStruct := range ants {
		for j, oneRoom := range path.Path {
			if oneRoom.IsStart || oneRoom.IsEnd {
				oneRoom.Ants = len(ants)
				oneAntStruct.CurrentRoom = oneRoom
				if oneRoom.IsEnd {
					endPrint := fmt.Sprintf("L%v-%v", oneAntStruct.Id, oneAntStruct.CurrentRoom.Name)
					result[j+i] = fmt.Sprintf("%s%s", result[j+i], endPrint)
				}
			} else if !oneRoom.IsStart && !oneRoom.IsEnd {
				oneRoom.Ants = +1
				oneAntStruct.CurrentRoom = oneRoom
				roomPrint := fmt.Sprintf("L%v-%v", oneAntStruct.Id, oneAntStruct.CurrentRoom.Name)
				result[j+i] = fmt.Sprintf("%s %s", result[j+i], roomPrint)
			}
		}
	}
	for k, trimResult := range result {
		result[k] = strings.TrimSpace(trimResult)
	}
	return deleteEmpty(result)
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

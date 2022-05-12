package ants

import (
	"fmt"
	"strconv"
	"strings"

	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

// AssignAnts assigns a path to each ant
func AssignAnts(ants []*structs.Ant, paths []*structs.PathStruct) []*structs.PathStruct {
	for _, ant := range ants {
		for i, path := range paths {
			if i+1 == len(paths) {
				if (len(path.Path) + len(path.Ants)) <= (len(paths[0].Path) + len(paths[0].Ants)) {
					path.Ants = append(path.Ants, ant)
					break
				}
			} else {
				if (len(path.Path) + len(path.Ants)) <= (len(paths[i+1].Path) + len(paths[i+1].Ants)) {
					path.Ants = append(path.Ants, ant)
					break
				}
			}
		}
	}
	return paths
}

// MoveAnts take the ants one by one and moves them through their paths
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

func GetSliceOfPathsNames(paths []*structs.PathStruct) []string {
	ret := []string{}
	for _, path := range paths {
		for _, pathInfo := range path.Path {
			ret = append(ret, pathInfo.Name)
		}
	}
	return ret
}

func GetSliceOfAntsNames(ants []*structs.Ant) []string {
	ret := []string{}
	for _, ant := range ants {
		ret = append(ret, strconv.Itoa(ant.Id))
	}
	return ret
}

func GetSliceOfRoomsNames(rooms []*structs.Room) []string {
	ret := []string{}
	for _, r := range rooms {
		ret = append(ret, r.Name)
	}
	return ret
}

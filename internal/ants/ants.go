package ants

import (
	"fmt"
	//"regexp"
	"strings"

	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

// fmt.Printf("Length of paths[i+1].Path: %+v paths[i+1].Ants: %+v\n", len(getSliceOfRoomNames(paths[i+1].Path)), getSliceOfAntsNames(paths[i+1].Ants))

// fmt.Printf("paths[i+1].Path: %+v paths[i+1].Ants: %+v\n", getSliceOfRoomNames(paths[i+1].Path), getSliceOfAntsNames(paths[i+1].Ants))

// fmt.Printf("path.Path: %+v path.Ants: %+v ant: %+v\n", getSliceOfRoomNames(path.Path), getSliceOfAntsNames(path.Ants), ant)

// func AssignAnts assigns a path to each ant
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

	numberofRooms := noOfMoves(ants, path)

	roomNumber := 1

	for i := 0; i < len(numberofRooms); i++ {
		for j := 0; j < numberofRooms[i]; j++ {
			//fmt.Println(numberofRooms[i], j)
			fmt.Println(result[roomNumber])
			roomNumber = roomNumber + 1
			if i < len(numberofRooms)-1 {
				i = i + 1
			}
		}
	}
	fmt.Println(result[len(result)-1])

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

func noOfMoves(ants []*structs.Ant, paths *structs.PathStruct) []int {
	//need to find a way to get number of paths - in this instance used 1
	lenPaths := 1
	rooms := []string{}
	rooms = append(rooms, "2", "3", "2", "1", "3", "2", "1", "3", "1")
	result := []int{}

	result = append(result, lenPaths)

	numberOfRooms := lenPaths
	var numberToAdd int

	for numberOfRooms < len(rooms) {
		noRooms := 0

		if numberToAdd < len(ants) {
			numberToAdd = result[len(result)-1] + lenPaths
			for i := numberOfRooms; i < numberOfRooms+numberToAdd; i++ {
				if i < numberOfRooms+numberToAdd {
					if rooms[i] == "1" || rooms[i] != rooms[i+1] {
						noRooms = noRooms + 1
					}
				}
			}
			numberOfRooms = numberOfRooms + noRooms
			result = append(result, noRooms)
		} else if numberToAdd == len(ants) {
			for i := numberOfRooms; i < numberOfRooms+numberToAdd; i++ {
				if i < numberOfRooms+numberToAdd {
					currentRoom := rooms[i]
					for r := 1; r < numberToAdd; r++ {
						if rooms[i+r] == currentRoom {
							break
						} else {
							noRooms = noRooms + 1
						}
					}
					numberToAdd = numberToAdd - 1
				}
			}
			result = append(result, noRooms)
			numberOfRooms = numberOfRooms + numberOfRooms
		} else {
			break
		}
	}

	//fmt.Println(noRooms)

	//loop through rooms from numberOfRooms i.e. numberOfRooms = 1 + len paths then go through the rooms
	//i.e. look at rooms[1] and rooms[2] do they equal the same if not add number 2 to result if both rooms are end room then add 3 (len ants)

	return result
}

// func getSliceOfAntNames(ants []*structs.Ant) []string {
// 	ret := []string{}
// 	for _, ant := range ants {
// 		ret = append(ret, strconv.Itoa(ant.Id))
// 	}
// 	return ret
// }

// func getSliceOfRoomNames(rooms []*structs.Room) []string {
// 	ret := []string{}
// 	for _, r := range rooms {
// 		ret = append(ret, r.Name)
// 	}
// 	return ret
// }

// func getSliceOfPathNames(paths *structs.PathStruct) []string {
// 	ret := []string{}
// 	for _, r := range paths.Path {
// 		ret = append(ret, r.Name)
// 	}
// 	return ret
// }
